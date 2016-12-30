package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	token := os.Getenv("SLACK_API_TOKEN")
	api := slack.New(token)
	// api.SetDebug(true)

	users, _ := api.GetUsers()
	leeroyID := ""
	for _, user := range users {
		if user.Name == "leeroy-jenkins" {
			leeroyID = user.ID
			fmt.Println("Found Leeroy!")
		}
	}

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Println("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				channel := ev.Channel
				userID := ev.User
				text := ev.Text
				fmt.Printf("#%s %s: %s\n", channel, userID, text)
				// Space odyssey fail
				if userID == leeroyID && strings.Contains(text, "Failure") {
					open.Run("https://www.youtube.com/watch?v=wpFQLw5_N2o")
				} else if userID == leeroyID && strings.Contains(text, "Success") {
					open.Run("https://www.youtube.com/watch?v=SBCw4_XgouA")
				}

			case *slack.InvalidAuthEvent:
				fmt.Println("Invalid credentials")
				return

			default:
				// Ignore other events..
				// fmt.Printf("Unexpected: %v\n", msg.Data)
			}
		}
	}
}
