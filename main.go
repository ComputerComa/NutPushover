package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gregdel/pushover"
	"github.com/joho/godotenv"
)

func main() {
	version := "1.0.0"
	if os.Args[1] == "-v" || os.Args[1] == "--v" || os.Args[1] == "--version" || os.Args[1] == "-version" {
		fmt.Println(version)
		os.Exit(0)
	}

	godotenv.Load("/etc/default/upsnotifyPushover")
	notificationType := os.Getenv("NOTIFYTYPE")
	upsName := os.Getenv("UPSNAME")
	notificationMessage := os.Args[1]
	apiKey := os.Getenv("PUSHOVER_API_KEY")
	userKey := os.Getenv("PUSHOVER_USER_KEY")
	app := pushover.New(apiKey)
	recipient := pushover.NewRecipient(userKey)
	title := fmt.Sprintf("UPS: %s (%s)", notificationType, upsName)
	message := pushover.NewMessageWithTitle(title, notificationMessage)
	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
	log.Println(response)

}
