package AdminCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
	"strconv"
	"strings"
	"time"
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
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "已将用户"+Parts[0]+"的会员设置为"+time.Unix(IntExpire, 0).Format("2006-01-02 15:04:05"))
	} else {
		_, err := Sends.SendTextMessage(UGid.MainId, UGid.MainType, "无权执行本命令")
		if err != nil {
			return
		}
	}
}
