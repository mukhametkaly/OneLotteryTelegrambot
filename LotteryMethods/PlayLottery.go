package LotteryMethods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)




func PlayLottery (lotID string, userID int) string {

	lottery, err := GetLotteryByID(lotID)
	if err != nil {
		return "Error"
	}
	if lottery == nil {
		return "Lottery not found"
	}

	if lottery.Raffler.UserID != userID {
		return ""
	}


	lottery, out := PlayLotteryRequest(lottery.LotteryID)
	if out != nil {
		return *out
	}
	return 	PrintLotteryWinner(*lottery)

}



func PlayLotteryRequest(lotteryID string) (*Lottery, *string) {

	url := "http://3.134.80.221/lottery/play/" + lotteryID
	output := "Error"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return nil, &output
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	lottery := &Lottery{}

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		fmt.Printf("runTransport %#v\n\n\n", string(respBody))
		output = string(respBody)
		return nil, &output
	}

	err = json.Unmarshal(respBody, lottery)

	if err != nil {
		fmt.Println(err.Error())
		return nil, &output
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