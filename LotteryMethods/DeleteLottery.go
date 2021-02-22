package LotteryMethods

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var toDelete[]*deleteLottery


type deleteLottery struct {
	UserID int
	Lottery *Lottery
}

func AppendToDelete(userID int)  {
	for _, i := range toDelete {
		if i.UserID == userID {
			return
		}
	}
	deleteLot := deleteLottery{
		UserID: userID,
	}

	toDelete = append(toDelete, &deleteLot)

}

func PrintUserLotteriesToDelete(userID int) string  {

	lotteries := GetLotteriesByRaffler(userID)
	if len(lotteries) == 0 {
		return "Sorry you dont have any lotteries"
	}

	str := "Please chose lottery which you need to delete\n"

	for _, i := range lotteries {
		str += `/` + "delete" + i.LotteryID + " - " + i.LotName + "\n"
	}
	return str

}

func SetLotteryToDelete(lotID string, userID int) string {

	for _, i := range toDelete {
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
			return DeleteLotteryRequest(i.Lottery)


		}
	}

	return ""


}



func DeleteLotteryRequest(lottery *Lottery) string {
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


	body := bytes.NewBuffer([]byte(""))

	url := "http://3.134.80.221/lottery/delete/" + lottery.LotteryID
	req, _ := http.NewRequest(http.MethodDelete, url, body)
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Content-Length", strconv.Itoa(0))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		return "Sorry error happend, lottery don't deleted"
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return "Lottery" + lottery.LotName + "successful deleted"
	} else {
		fmt.Printf("runTransport %#v\n\n\n", string(respBody))
		return string(respBody)
	}

}