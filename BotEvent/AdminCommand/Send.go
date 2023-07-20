package AdminCommand

import (
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
	"strings"
)

func Send(Content string, UGid Type.Id) {
	Parts := strings.SplitN(Content, "|", 2)
	_, err := Sends.SendTextMessage(Parts[0], "user", Parts[1])
	if err != nil {
		log.Println(err)
		return
	}
	_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "发送成功")
	if err != nil {
		log.Println(err)
		return
	}
}
