package LotteryMethods

import (
	"encoding/json"
	"fmt"
	"github.com/mukhametkaly/OneLotteryTelegrambot/Bot"
	"io/ioutil"
	"net/http"
	"strconv"
)

var toPlay[]*playLottery


type playLottery struct {
	UserID int
	Lottery *Lottery
}

func AppendToPlay(userID int)  {
	for _, i := range toPlay {
		if i.UserID == userID {
			return
		}
	}
	playLot := playLottery{
		UserID: userID,
	}

	toPlay = append(toPlay, &playLot)

}

func PrintUserLotteriesToPlayLottery(userID int) string  {

	lotteries := GetLotteriesByRaffler(userID)
	if len(lotteries) == 0 {
		return "Sorry you dont have any lotteries"
	}

	str := "Please chose lottery which you need to play\n"

	for _, i := range lotteries {
		str += `/` + "play" + i.LotteryID + " - " + i.LotName + "\n"
	}
	return str

}

func SetLotteryToPlay (lotID string, userID int) string {

	for _, i := range toPlay {
		if i.UserID == userID {
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

			i.Lottery = lottery

			lottery, err = PlayLotteryRequest(i.Lottery.LotteryID)
			if err != nil {
				return "Something went wrong"
			}
			return 	Bot.PrintLotteryWinner(*lottery)


		}
	}

	return ""


}



func PlayLotteryRequest(lotteryID string) (*Lottery, error) {

	url := "http://3.134.80.221/lottery/play/" + lotteryID
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	lottery := &Lottery{}

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		fmt.Printf("runTransport %#v\n\n\n", string(respBody))
	}

	err = json.Unmarshal(respBody, lottery)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}


	return lottery, nil

}




func EnjoyLotteryRequest(lotteryID string, userID int, username string) string {
	idParam := strconv.Itoa(userID)

	url := "http://3.134.80.221/lottery/" + lotteryID + "/newplayer/" + idParam +  "/username/" + username
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return "Sorry error happend"
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)



	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		fmt.Printf("runTransport %#v\n\n\n", string(respBody))
	}



	return string(respBody)

}