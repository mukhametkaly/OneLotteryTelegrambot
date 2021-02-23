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
				message := LotteryMethods.PrintUserLotteries(userID, "update")
				bot.Send(NewReplyMessage(update.Message, message))

			case "play":

				message := LotteryMethods.PrintUserLotteries(userID, "play")
				bot.Send(NewReplyMessage(update.Message, message))

			case "delete":
				message := LotteryMethods.PrintUserLotteries(userID, "delete")
				bot.Send(NewReplyMessage(update.Message, message))

			case "info":
				message := LotteryMethods.PrintUserLotteries(userID, "info")
				bot.Send(NewReplyMessage(update.Message, message))

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

			case "timer":
				LotteryMethods.AppendToTimers(userID)
				message := LotteryMethods.PrintUserLotteries(userID, "timer")
				bot.Send(NewReplyMessage(update.Message, message))

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
						result := LotteryMethods.DeleteLottery(*lotID, userID)
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
						result := LotteryMethods.PlayLottery(*lotID, userID)
						if result == "" {

						} else {
							bot.Send(NewReplyMessage(update.Message, result))
						}

					case "info":
						result := LotteryMethods.GetLotteryInfo(*lotID, userID)
						if result == "" {

						} else {

							bot.Send(NewReplyMessage(update.Message, result))
						}

					case "timer":
						result := LotteryMethods.SetLotteryToTimer(*lotID, userID)
						if result == "" {
						} else {
							bot.Send(NewReplyMessage(update.Message, result))
						}

					}
				}
			}
		} else {

			 if update.Message.ReplyToMessage != nil {
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

					} else if update.Message.ReplyToMessage.Text == "How many hours do you want to set the timer for?" {
						timer, err := strconv.Atoi(update.Message.Text)
						if err != nil {
							bot.Send(NewReplyMessage(update.Message, "Incorrect please write number in correct format"))
						} else {
							resultID := LotteryMethods.GetTimer(userID)
							go PlayLotteryAfterTimer(timer, update.Message.Chat.ID, userID, resultID.LotID, bot  )

							bot.Send(NewReplyMessage(update.Message, "OKQ! Timer started" ))
						}


					}

				}
			}
		}
	}
}

