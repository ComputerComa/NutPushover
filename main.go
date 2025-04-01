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
	debug := false
	version := "1.0.2"
	if os.Args[1] == "-v" || os.Args[1] == "--v" || os.Args[1] == "--version" || os.Args[1] == "-version" {
		fmt.Println(version)
		os.Exit(0)
	}

	godotenv.Load("/etc/default/upsnotifyPushover")

	if os.Getenv("DEBUG") == "True" || os.Getenv("DEBUG") == "true" {
		debug = true
	}

	hostname, herr := os.Hostname()

	if herr != nil {
		log.Panic(herr)
	}
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
	fnotifmsg := fmt.Sprintf("%s -- %s", hostname, notificationMessage)
	message := pushover.NewMessageWithTitle(fnotifmsg, title)

	if len(sound) != 0 {
		message.Sound = sound
	}

	if len(priority) != 0 {
		tmp_pri := parsePriority(priority)
		if tmp_pri == pushover.PriorityEmergency {
			message.Retry = retry
			message.Expire = expire

		}
		message.Priority = tmp_pri
	}

	if debug {

		fmt.Printf("Sending Pushover notification \nNotification Type - %s \nMessage Priority - %s \nMessage Sound - %s \nUPS Name - %s \nMessage - %s \nTitle - %s \n", notificationType, priority, sound, upsName, fnotifmsg, title)
	}
	response, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
	log.Println(response)

}
