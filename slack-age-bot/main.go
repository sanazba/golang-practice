package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3658282092596-3680009889152-XZsfPc2arF0MibOV9aZJMwkA")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03KZJXT21W-3658284828660-469b5bb91a822d2766b207f1919b26570ff2bc445a8218661627f5bd29ed0c62")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Example:     "my yob is 2020",
		Handler: func(botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("errors")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age) // return the resulting string instead of printing to stdout
			response.Reply(r)
		},
	},
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
