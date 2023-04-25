package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type UnsplashPhoto struct {
	Urls struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}

func GetRandomPhoto() (string, error) {
	url := "https://api.unsplash.com/photos/random?client_id=" + os.Getenv("UNSPLASH_ACCESS_KEY") // random photo

	resp, err := http.Get(url) // trying get request
	if err != nil {
		return "", err
	}

	fmt.Println(resp.Body)

	defer resp.Body.Close() // will be closed at the end

	var photo UnsplashPhoto // variable

	err = json.NewDecoder(resp.Body).Decode(&photo) // store resp.Body to photo variable
	if err != nil {
		return "", err
	}

	return photo.Urls.Regular, nil
}

func main() {
	err := godotenv.Load("../.env") // necessary for loading from .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0) // update bot configuration each time
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig) // getting updates which is stored in channels

	countCh := make(chan int)
	count := 0

	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}

			log.Printf("[%s]%s", update.Message.From.UserName, update.Message.Text)

			// check if user sends "/image" or "image" command
			if update.Message.IsCommand() && update.Message.Command() == "image" || update.Message.Text == "image" {
				countCh <- 1

				photo, err := GetRandomPhoto()
				if err != nil {
					log.Println(err)
				}

				file := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(photo)) // instead of NewMessage

				bot.Send(file)
			}
		}
	}()

	go func() {
		for {
			<-countCh
			count++

			log.Printf("Count : %v", count)
		}
	}()

	select {}
}
