package AdminCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
	"strconv"
	"strings"
)

func SetUserPremium(Content string, UGid Type.Id) {
	IsAmin, _ := SQLite.IsAmin(UGid.User)
	if IsAmin {
		Parts := strings.SplitN(Content, "|", 2)

		IntExpire, err := strconv.ParseInt(Parts[1], 10, 64)

		err = SQLite.UpdateUserPremium(Parts[0], true, IntExpire)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		_, err := Sends.SendTextMessage(UGid.MainId, UGid.MainType, "无权执行本命令")
		if err != nil {
			return
		}
	}
}
