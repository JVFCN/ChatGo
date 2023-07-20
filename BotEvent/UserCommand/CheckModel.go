package UserCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
)

func Model(Type string, UGid Type.Id) {
	_, err := Sends.SendTextMessage(UGid.MainId, Type, UGid.Name+"的模型是:"+SQLite.GetUserModel(UGid.User))
	if err != nil {
		log.Println(err)
	}
}
