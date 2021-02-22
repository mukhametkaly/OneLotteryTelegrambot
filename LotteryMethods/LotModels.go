package LotteryMethods

import (
	"time"
)


type Lottery struct {
	LotteryID   string    `json:"lottery_id"`
	LotName     string    `json:"lot_name"`
	Raffler     User      `json:"raffler"`
	Winner      *User     `json:"winner, omitempty"`
	Starttime   time.Time `json:"startime"`
	Prize       string    `json:"prize"`
	TextMessage string    `json:"text_message"`
	PlayerIDs   []User    `json:"player_ids"`
}

type User struct {
	Username string `json:"username"`
	UserID int `json:"user_id,omitempty"`
}



