package main

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"sync"
)

type UnsplashPhoto struct {
	Urls struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}

func GetRandomPhoto() (string, error) {
	url := "https://api.unsplash.com/photos/random?client_id=" + os.Getenv("UNSPLASH_ACCESS_KEY")
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var photo UnsplashPhoto
	err = json.NewDecoder(resp.Body).Decode(&photo)
	if err != nil {
		return "", err
	}
	return photo.Urls.Regular, nil
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	var mutex sync.Mutex
	count := 0

	go func() {
		for update := range updates {
			if update.Message != nil {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				if update.Message.IsCommand() && update.Message.Command() == "image" || update.Message.Text == "image" {
					mutex.Lock()
					count++
					log.Printf("Count: %v", count)
					photo, err := GetRandomPhoto()
					if err != nil {
						log.Println(err)
					}
					file := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(photo))
					bot.Send(file)
					mutex.Unlock()
				}
			}
		}
	}()
	select {}
}
