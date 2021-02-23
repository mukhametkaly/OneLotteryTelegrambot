package Bot

import (
	"github.com/mukhametkaly/OneLotteryTelegrambot/LotteryMethods"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"time"
)

func NewReplyMessage(message *tgbotapi.Message, text string) *tgbotapi.MessageConfig {
	return &tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           message.Chat.ID,
			ReplyToMessageID: message.MessageID,
		},
		Text:                  text,
		ParseMode: 			   tgbotapi.ModeHTML,
		DisableWebPagePreview: false,
	}
}

func NewMarkdownMessage(ChatID int64, text string) *tgbotapi.MessageConfig {
	return &tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           ChatID,
		},
		Text:                  text,
		ParseMode: 			   tgbotapi.ModeHTML,
		DisableWebPagePreview: false,
	}
}


func SeparateCommandAndLotID(text string) (Command, lotID *string) {
	command := text
	commandSize := len(command)

	if commandSize > 6 {

		if command[0:6]  == "update" {

			resultID := command[6:commandSize]
			resultCommand := "update"
			return &resultCommand, &resultID

		} else if command[0:6]  == "delete" {

			resultID := command[6:commandSize]
			resultCommand := "delete"
			return &resultCommand, &resultID

		}else if command[0:5]  == "enjoy" {

			resultID := command[5:commandSize]
			resultCommand := "enjoy"
			return &resultCommand, &resultID

		}else if command[0:4] == "play" {
			resultID := command[4:commandSize]
			resultCommand := "play"
			return &resultCommand, &resultID
		}else if command[0:4] == "info" {
			resultID := command[4:commandSize]
			resultCommand := "info"
			return &resultCommand, &resultID
		} else if command[0:5] == "timer" {
			resultID := command[5:commandSize]
			resultCommand := "timer"
			return &resultCommand, &resultID
		}

	}

	return nil, nil

}


func PlayLotteryAfterTimer(timer int, chatID int64, userID int, lotID string, bot *tgbotapi.BotAPI)  {
	time.Sleep(time.Second * time.Duration(timer))
	str := LotteryMethods.PlayLottery( lotID, userID)
	bot.Send(NewMarkdownMessage(chatID, str))

}
