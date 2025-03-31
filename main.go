package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gregdel/pushover"
	"github.com/joho/godotenv"
)

func parsePriority(p string) int {
	out_prio := pushover.PriorityNormal
	switch p {
	case "-2":
		out_prio = pushover.PriorityLowest
	case "-1":
		out_prio = pushover.PriorityLow
	case "0":
		out_prio = pushover.PriorityNormal
	case "1":
		out_prio = pushover.PriorityHigh
	case "2":
		out_prio = pushover.PriorityEmergency
	default:
		out_prio = pushover.PriorityNormal
	}
	return out_prio

}

func main() {
	version := "0.0.1"
	if os.Args[1] == "-v" || os.Args[1] == "--v" || os.Args[1] == "--version" || os.Args[1] == "-version" {
		fmt.Println(version)
		os.Exit(0)
	}

	godotenv.Load("/etc/default/upsnotifyPushover")
	notificationType := os.Getenv("NOTIFYTYPE")
	priority := os.Getenv("PUSHOVER_Priority")
	sound := os.Getenv("PUSHOVER_Sound")
	retry := 60 * time.Second
	expire := 30 * time.Minute
	upsName := os.Getenv("UPSNAME")
	notificationMessage := os.Args[1]
	apiKey := os.Getenv("PUSHOVER_API_KEY")
	userKey := os.Getenv("PUSHOVER_USER_KEY")
	app := pushover.New(apiKey)
	recipient := pushover.NewRecipient(userKey)
	title := fmt.Sprintf("UPS: %s (%s)", notificationType, upsName)
	message := pushover.NewMessageWithTitle(title, notificationMessage)

	if len(sound) != 0 {
		message.Sound = sound
	}

	if len(priority) != 0 {
		tmp_pri := parsePriority(priority)
		if tmp_pri == pushover.PriorityEmergency {
			message.Retry = retry
			message.Expire = expire

		}
	}

	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
	log.Println(response)

}
