package UserCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
	"time"
)

func Premium(Type string, UGid Type.Id) {
	ExpireTime := time.Unix(SQLite.GetUserPremiumExpire(UGid.User), 0).Unix()
	if ExpireTime <= time.Now().Unix() {
		_, err := Sends.SendTextMessage(UGid.MainId, Type, UGid.Name+"的会员已经到期了")
		if err != nil {
			log.Println(err)
		}
		return
	} else {
		_, err := Sends.SendTextMessage(UGid.MainId, Type, UGid.Name+"的会员到期时间为:"+time.Unix(ExpireTime, 0).Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Println(err)
		}
	}
}
