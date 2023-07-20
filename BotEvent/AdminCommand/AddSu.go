package AdminCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
)

func AddSu(UGid Type.Id, UserId string) {
	IsAmin, _ := SQLite.IsAmin(UGid.User)
	if IsAmin {
		err := SQLite.UpdateUserAdmin(UserId, true)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		_, err := Sends.SendTextMessage(UGid.MainId, UGid.MainType, "无权执行本命令")
		if err != nil {
			log.Println(err)
			return
		}
	}
}
