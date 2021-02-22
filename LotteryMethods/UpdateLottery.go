package LotteryMethods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

var toUpdate []*updateLottery


type updateLottery struct {
	UserID int
	UpdateParam string
	Lottery *Lottery
}

func AppendToUpdate(userID int)  {
	for _, i := range toUpdate {
		if i.UserID == userID {
			return
		}
	}
	updLot := updateLottery{
		UserID: userID,
	}

	toUpdate = append(toUpdate, &updLot)

}

func PrintUserLotteriesToUpdate(userID int) string  {

	lotteries := GetLotteriesByRaffler(userID)
	if len(lotteries) == 0 {
		return "Sorry you dont have any lotteries"
	}

	str := "Please chose lottery which you need to update\n"

	for _, i := range lotteries {
		str += `/` + "update" + i.LotteryID + " - " + i.LotName + "\n"
	}
	return str

}

func SetLotteryToUpdate(lotID string, userID int) string {

	for _, i := range toUpdate {
		if i.UserID == userID {
			lottery, err := GetLotteryByID(lotID)
			if err != nil {
				return "Something went wrong, error to update"
			}
			if lottery == nil {
				return "Something went wrong, error to update"
			}

			if lottery.Raffler.UserID != userID {
				return ""
			}

			i.Lottery = lottery

			return "Specify which parameter you want to update \n" +
				`/` + "text to edit Lottery text \n" +
				`/` + "prize to edit Lottery prize \n" +
				`/` + "name to edit Lottery name \n"

		}
	}

	return ""


}


func SetUpdateParam (updText string, userID int) string {
	for _, i := range toUpdate {
		if i.UserID == userID {
			if i.Lottery == nil {
				return ""
			}
			i.UpdateParam = updText
			return "OK"
		}
	}
	return ""
}


func UpdateLottery(updText string, userID int) string  {
	for _, i := range toUpdate {
		if i.UserID == userID {

			switch i.UpdateParam {
				case "prize":
					i.Lottery.Prize = updText
				case "text":
					i.Lottery.TextMessage = updText
				case "name":
					i.Lottery.LotName = updText
			}
			return UpdateLotteryRequest(i.Lottery)
		}
	}
	return "Nothing to update"
}


func UpdateLotteryRequest(lottery *Lottery) string {
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

	url := "http://3.134.80.221/lottery/update"
	req, _ := http.NewRequest(http.MethodPut, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		return "Sorry error happend, lottery don't updated"
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return "Lottery" + lottery.LotName + "successful updated"
	} else {
		fmt.Printf("runTransport %#v\n\n\n", string(respBody))
		return string(respBody)
	}

}