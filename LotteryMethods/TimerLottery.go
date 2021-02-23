package LotteryMethods


var toTimer []*TimerLottery


type TimerLottery struct {
	UserID int
	LotID string
	isActive bool
}

func GetTimer(userID int) *TimerLottery {
	for _, i := range toTimer {
		if i.UserID == userID {
			return i
		}
	}
	return nil
}

func AppendToTimers(userID int)  {
	for _, i := range toTimer {
		if i.UserID == userID {
			return
		}
	}
	timerLot := TimerLottery{
		UserID: userID,
	}

	toTimer = append(toTimer, &timerLot)
}


func SetLotteryToTimer(lotID string, userID int) string {

	for _, i := range toTimer {
		if i.UserID == userID {
			lottery, err := GetLotteryByID(lotID)
			if err != nil {
				return "Something went wrong, error to start timer"
			}
			if lottery == nil {
				return "Something went wrong, error to start timer"
			}

			if lottery.Raffler.UserID != userID {
				return ""
			}

			i.LotID = lotID

			return "How many hours do you want to set the timer for?"
		}
	}
	return ""

}

