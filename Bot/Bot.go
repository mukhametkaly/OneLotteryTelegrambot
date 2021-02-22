package Bot

import (
	"github.com/mukhametkaly/OneLotteryTelegrambot/LotteryMethods"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func NewReplyMessage(message *tgbotapi.Message, text string) *tgbotapi.MessageConfig {
	return &tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           message.Chat.ID,
			ReplyToMessageID: message.MessageID,
		},
		Text:                  text,
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
		}

	}

	return nil, nil

}

func PrintLotteryMessage (lottery LotteryMethods.Lottery) string {
	return `*` + lottery.LotName + `*` + "\n\n" +
		lottery.TextMessage + "\n\n" +
		"Prize: " + lottery.Prize + "\n\n"+
		"To enjoy this lottery: \n" + `/enjoy` + lottery.LotteryID

}


func PrintLotteryWinner (lottery LotteryMethods.Lottery) string {
	//str := strconv.Itoa(lottery.Winner.UserID)
	return `**` + lottery.LotName + `**` + "\n\n" +
		lottery.TextMessage + "\n\n" +
		"Prize: " + lottery.Prize + "\n\n"+
		"CONGRATULATIONS TO THE WINNER: \n" +
		"@" + lottery.Winner.Username
		//"[" + lottery.Winner.Username + "](tg://user?id=" + str + ")"


}