package LotteryMethods

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetLotteriesByRaffler(UserID int) []Lottery {

	idparam := strconv.Itoa(UserID)
	url := "http://3.134.80.221/lottery/raffler/" + idparam

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return nil
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	var lottery []Lottery

	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		fmt.Printf("runTransport %#v\n\n\n", string(respBody))
	}

	err = json.Unmarshal(respBody, &lottery)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}


	return lottery

}

func GetLotteryByID(id string) (*Lottery, error) {

	url := "http://3.134.80.221/lottery/" + id
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