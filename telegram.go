package main

import "github.com/mukhametkaly/OneLotteryTelegrambot/Bot"


const (
	BotToken   = "1659894835:AAFMGTrQgxWRc8Qh8rlGRiziuZ1mxMBs7iA"
	WebhookURL = "https://3d009753deb8.ngrok.io"
)


func main() {

	myBot := Bot.NewBot(BotToken)
	myBot.Run(WebhookURL)

}