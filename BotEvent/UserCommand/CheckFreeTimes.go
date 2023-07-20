package UserCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
	"strconv"
)

func CheckFreeTimes(UGid Type.Id) {
	FreeTimes := SQLite.GetUserFreeTimes(UGid.User)

	_, err := Sends.SendTextMessage(UGid.MainId, UGid.MainType, "您的免费次数还剩"+strconv.Itoa(FreeTimes)+"次")
	if err != nil {
		log.Println(err)
		return
	}
}
