package LotteryMethods

import (
	"strconv"
	"time"
)

func FormatTime(time time.Time) string {
	result := ""
	d, m, y := time.Date()
	minute := strconv.Itoa(time.Minute())
	if time.Minute() < 10 {
		minute = "0" + minute
	}
	if time.Hour() >= 18 {
		result += strconv.Itoa(d) + ` ` + m.String() + ` ` + strconv.Itoa(y+1) + " hour: "
		result += strconv.Itoa(time.Hour() - 18) + `:` + minute
		return result
	}
	result += strconv.Itoa(d) + ` ` + m.String() + ` ` + strconv.Itoa(y)
	result += strconv.Itoa(time.Hour() + 6) + `:` + minute
	return result

}

func PrintUserLotteries(userID int, command string) string  {

	lotteries := GetLotteriesByRaffler(userID)
	if len(lotteries) == 0 {
		return "Sorry you dont have any lotteries"
	}
	str := ""
	if command == "info" {
		str = "Please chose lottery which you need to get " + command + "\n"
	} else if  command == "timer" {
		str = "Please chose lottery which you need to set" + command + "\n"
	} else {
		str = "Please chose lottery which you need to" + command + "\n"

	}

	for _, i := range lotteries {
		str += `/` + command + i.LotteryID + " - " + i.LotName + "\n"
	}
	return str

}
func PrintLotteryMessage (lottery Lottery) string {
	return "<b>" + lottery.LotName + "</b>" + " \n\n" +
		lottery.TextMessage + "\n\n" +
		"Prize: " + lottery.Prize + "\n\n"+
		"Start time: " + FormatTime(lottery.Starttime) + "\n\n"+
		"To enjoy this lottery: \n" + `/enjoy` + lottery.LotteryID
}

func PrintLotteryWinner (lottery Lottery) string {

	str := strconv.Itoa(lottery.Winner.UserID)
	return "<b>" + lottery.LotName + "</b>" + " \n\n" +
		lottery.TextMessage + "\n\n" +
		"Prize: " + lottery.Prize + "\n\n"+
		"Start time: " + FormatTime(lottery.Starttime) + "\n\n"+
		"CONGRATULATIONS TO THE WINNER: \n" +
		"<a href=" + `"` + "tg://user?id=" + str + `"` + "><b>WINNER</b></a>"


		//"[!!WINNER!!](tg://user?id=" + str + ")"


}