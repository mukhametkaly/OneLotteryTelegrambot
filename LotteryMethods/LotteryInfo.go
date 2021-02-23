package LotteryMethods




func GetLotteryInfo(lotID string, userID int) string {

	lottery, err := GetLotteryByID(lotID)
	if err != nil {
		return "Something went wrong, error to delete"
	}
	if lottery == nil {
		return "Lottery not found"
	}

	if lottery.Raffler.UserID != userID {
		return ""
	}

	if lottery.Winner != nil {
		return PrintLotteryWinner(*lottery)
	}

	return 	PrintLotteryMessage(*lottery)


}

