package LotteryMethods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mukhametkaly/OneLotteryTelegrambot/Bot"

	//tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

var BeingCreatedLotteryCollection []*Lottery


func AppendEmptyLottery(userID int, userName string)  {
	for _, i := range BeingCreatedLotteryCollection {
		if i.Raffler.UserID ==  userID {
			return
		}
	}
	emptyLot := Lottery{}
	emptyLot.Raffler = User{
		UserID: userID,
		Username: userName,
	}
	BeingCreatedLotteryCollection = append(BeingCreatedLotteryCollection, &emptyLot)
}

func SetLotteryName( userID int, LotName string) bool {
	for _, i := range BeingCreatedLotteryCollection {
		if i.Raffler.UserID ==  userID {
			i.LotName = LotName
			return true
		}
	}
	return false
}

func SetLotteryPrize ( userID int, prize string) bool {
	for _, i := range BeingCreatedLotteryCollection {
		if i.Raffler.UserID ==  userID {
			i.Prize = prize
			return true
		}
	}
	return false
}

func SetLotteryText( userID int, text string) bool {
	for _, i := range BeingCreatedLotteryCollection {
		if i.Raffler.UserID ==  userID {
			i.TextMessage = text
			return true
		}
	}
	return false
}

//func CreateLotteryWithTimer(timer int, chatID uint64, userID int, bot *tgbotapi.BotAPI)  {
//	lottery := Lottery{}
//	for _, i := range BeingCreatedLotteryCollection {
//		if i.Raffler.UserID ==  userID {
//			lottery = i
//		}
//	}
//
//	err := CreateLotteryRequest(lottery)
//	if err != nil {
//		panic(err)
//	}
//}
//func LotteryTimer(timer int64, chatID uint64, bot *tgbotapi.BotAPI, lotID string)  {
//
//	time.Sleep(time.Duration(timer))
//
//}

func CreateLottery(userID int) string {
	lottery := Lottery{}
	for _, i := range BeingCreatedLotteryCollection {
		if i.Raffler.UserID ==  userID {
			lottery = *i
		}
	}
	RespLottery := CreateLotteryRequest(lottery)
	if RespLottery == nil {
		return ""
	}
	return Bot.PrintLotteryMessage(*RespLottery)

}


func CreateLotteryRequest(lottery Lottery) *Lottery {

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	data, err := json.Marshal(lottery)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(data)

	url := "http://3.134.80.221/lottery/create"
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		return nil
	}


	defer resp.Body.Close() // важный пункт!
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(respBody, &lottery)
	if err != nil {
		return nil
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return &lottery
	} else {
		fmt.Printf("runTransport %#v\n\n\n", string(respBody))
		return nil
	}


}
