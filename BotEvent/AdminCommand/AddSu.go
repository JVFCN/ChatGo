package AdminCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
)

func AddSu(UGid Type.Id, UserId string) {
	IsAmin, _ := SQLite.IsAmin(UGid.User)

	Id := ""
	if UserId == "me" {
		Id = UGid.User
	} else {
		Id = UserId
	}

	if IsAmin || UGid.User == "3161064" {
		err := SQLite.UpdateUserAdmin(Id, true)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "已将用户"+UGid.Name+"("+Id+")"+"设置为管理员")
	} else {
		_, err := Sends.SendTextMessage(UGid.MainId, UGid.MainType, "无权执行本命令")
		if err != nil {
			log.Println(err)
			return
		}
	}
}
