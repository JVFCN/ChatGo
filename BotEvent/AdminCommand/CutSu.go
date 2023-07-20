package AdminCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
)

func CutSu(UGid Type.Id, UserId string) {
	IsAmin, err := SQLite.IsAmin(UGid.User)
	if err != nil {
		return
	}
	if IsAmin {
		err := SQLite.UpdateUserAdmin(UserId, false)
		if err != nil {
			return
		}
	} else {
		_, err := Sends.SendTextMessage(UGid.MainId, UGid.MainType, "无权执行本命令")
		if err != nil {
			return
		}
	}
}
