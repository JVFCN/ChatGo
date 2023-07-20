package AdminCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
	"strconv"
	"strings"
)

func SetFree(UGid Type.Id, Content string) {
	IsAdmin, err := SQLite.IsAmin(UGid.User)
	if err != nil {
		log.Println(err)
		return
	}

	if IsAdmin {
		Parts := strings.SplitN(Content, "|", 2)

		FreeTimes, err := strconv.Atoi(Parts[1])
		if err != nil {
			log.Println(err)
			return
		}
		if Parts[0] == "me" {
			Parts[0] = UGid.User
		}

		err = SQLite.UpdateUserFreeTimes(Parts[0], FreeTimes)
		if err != nil {
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "已将用户"+Parts[0]+"的免费次数设置为"+Parts[1])
	} else {
		_, err := Sends.SendTextMessage(UGid.MainId, UGid.MainType, "无权执行本命令")
		if err != nil {
			return
		}
	}
}
