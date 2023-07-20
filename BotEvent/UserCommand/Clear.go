package UserCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
)

func Clear(Type string, UGid Type.Id) {
	err := SQLite.ClearUserContext(UGid.User)
	_, err = Sends.SendTextMessage(UGid.MainId, Type, "已为"+UGid.Name+"清除上下文")
	if err != nil {
		log.Println(err)
	}
}
