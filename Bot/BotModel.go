package Bot

import (
	"fmt"
	"github.com/mukhametkaly/OneLotteryTelegrambot/LotteryMethods"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"net/http"
	"strconv"
)

type MyBot struct {
	Bot *tgbotapi.BotAPI

}

func NewBot(token string)  *MyBot {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	myBot := MyBot{
		Bot: bot,
	}

	return &myBot
}



func (bt *MyBot) Run(WebhookURL string) {
	bot := bt.Bot

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook("/")

	go http.ListenAndServe(":8080", nil)

	fmt.Println("start listen :8080")

	for update := range updates {
		userID := update.Message.From.ID

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "create":
				LotteryMethods.AppendEmptyLottery(userID, update.Message.From.UserName)
				bot.Send(NewReplyMessage(update.Message, "Please write a name of your lottery"))

			case "update":
				LotteryMethods.AppendToUpdate(userID)
				message := LotteryMethods.PrintUserLotteriesToUpdate(userID)
				bot.Send(NewReplyMessage(update.Message, message))

			case "play":

				LotteryMethods.AppendToPlay(userID)
				message := LotteryMethods.PrintUserLotteriesToPlayLottery(userID)
				bot.Send(NewReplyMessage(update.Message, message))

			case "delete":
				LotteryMethods.AppendToDelete(userID)
				message := LotteryMethods.PrintUserLotteriesToDelete(userID)
				bot.Send(NewReplyMessage(update.Message, message))

			case "info":


			case "prize":
				result := LotteryMethods.SetUpdateParam("prize", userID)
				if result == "OK" {
					bot.Send(NewReplyMessage(update.Message, "Please write the update data"))
				}
			case "text":
				result := LotteryMethods.SetUpdateParam("text", userID)
				if result == "OK" {
					bot.Send(NewReplyMessage(update.Message, "Please write the update data"))
				}
			case "name":
				result := LotteryMethods.SetUpdateParam("name", userID)
				if result == "OK" {
					bot.Send(NewReplyMessage(update.Message, "Please write the update data"))
				}

			default:

				command, lotID := SeparateCommandAndLotID(update.Message.Command())
				if command == nil || lotID == nil {

				} else {
					switch *command {
					case "update":

						result := LotteryMethods.SetLotteryToUpdate(*lotID, userID)
						if result == "" {

						} else {
							bot.Send(NewReplyMessage(update.Message, result))
						}

					case "delete":
						result := LotteryMethods.SetLotteryToDelete(*lotID, userID)
						if result == "" {

						} else {
							bot.Send(NewReplyMessage(update.Message, result))
						}
					case "enjoy":
						var result string
						if update.Message.From.UserName == "" {
							username  := strconv.Itoa(userID)
							result = LotteryMethods.EnjoyLotteryRequest(*lotID, userID, username)
						} else {
							result = LotteryMethods.EnjoyLotteryRequest(*lotID, userID, update.Message.From.UserName)
						}
						if result == "" {

						} else {
							bot.Send(NewReplyMessage(update.Message, result))
						}
					case "play":
						result := LotteryMethods.SetLotteryToPlay(*lotID, userID)
						if result == "" {

						} else {
							bot.Send(NewReplyMessage(update.Message, result))
						}
					}
				}
			}
		} else {

			if update.Message.Text == "privet" && update.Message.From.UserName == "mukhametkaly" {
				bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"krasava",
				))
			} else if update.Message.ReplyToMessage != nil {
				if update.Message.ReplyToMessage.From.IsBot {

					if update.Message.ReplyToMessage.Text == "Please write a name of your lottery" {
						LotteryMethods.SetLotteryName(userID, update.Message.Text)
						bot.Send(NewReplyMessage(update.Message, "Please write what a prize in your lottery"))

					} else if update.Message.ReplyToMessage.Text == "Please write what a prize in your lottery" {
						LotteryMethods.SetLotteryPrize(userID, update.Message.Text)
						bot.Send(NewReplyMessage(update.Message, "Please write what a text in your lottery"))

					} else if update.Message.ReplyToMessage.Text == "Please write what a text in your lottery" {
						LotteryMethods.SetLotteryText(userID, update.Message.Text)
						message := LotteryMethods.CreateLottery(userID)
						bot.Send(NewReplyMessage(update.Message, message))

					} else if update.Message.ReplyToMessage.Text == "Please write the update data" {

						message := LotteryMethods.UpdateLottery(update.Message.Text, userID)
						bot.Send(NewReplyMessage(update.Message, message))

					}

				}
			}
		}
	}
}

