package BotEvent

import (
	"ChatGPTBot/BotEvent/AdminCommand"
	"ChatGPTBot/BotEvent/UserCommand"
	"ChatGPTBot/OpenAI"
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"github.com/yhchat/bot-go-sdk/subscription"
	"log"
	"regexp"
	"strings"
)

func Normal(event subscription.MessageEvent) {
	SendContent := event.Message.Content
	ChatType := event.Chat.ChatType
	UGid := Type.Id{}

	if ChatType == "bot" {
		ChatType = "user"
	} else {
		ChatType = "group"
	}

	if ChatType == "group" {
		UGid = Type.Id{
			MainId:   event.Chat.ChatId,
			MainType: event.Chat.ChatType,
			User:     event.Sender.SenderId,
			Group:    event.Chat.ChatId,
			Name:     event.Sender.SenderNickname,
		}
	} else {
		UGid = Type.Id{
			MainId:   event.Sender.SenderId,
			MainType: event.Sender.SenderType,
			User:     event.Sender.SenderId,
			Group:    event.Chat.ChatId,
			Name:     event.Sender.SenderNickname,
		}
	}

	err := SQLite.AddUser(UGid.User)
	if err != nil {
		log.Println(err)
		return
	}

	for _, Value := range SendContent {
		if StringValue, Ok := Value.(string); Ok {
			FirstChar := StringValue[0]
			if FirstChar == '.' {
				log.Println("[ChatGo]"+UGid.Name+"使用了用户命令:", Value.(string))
				CommandName, CommandContent := PartUserCommand(Value.(string))
				err := RunCommand(CommandName, CommandContent, ChatType, UGid)
				if err != nil {
					log.Println(err)
					return
				}
			} else if FirstChar == '!' {
				log.Println("[ChatGo]"+UGid.Name+"使用了管理员命令:", Value.(string))
				CommandName, CommandContent := PartAdminCommand(Value.(string))
				err := RunCommand(CommandName, CommandContent, ChatType, UGid)
				if err != nil {
					log.Println(err)
					return
				}
			} else if StringValue == "speak" {
				recentData := OpenAI.RecentDataList1.GetRecentData()
				text := ""
				for _, item := range recentData {
					text += item.Timestamp.Format("2006-01-02 15:04:05") + " " + item.Value + "\n"
				}
				_, err := Sends.SendTextMessage("3161064", "user", text)
				if err != nil {
					log.Println(err)
					return
				}
			} else {
				if ChatType == "group" && FindAt(StringValue) == false {
					return
				}
				Response, err := Sends.SendMarkdownMessage(UGid.MainId, UGid.MainType, "正在思考中...")
				if err != nil {
					return
				}
				MessageId := Response.Data.(map[string]interface{})["messageInfo"].(map[string]interface{})["msgId"].(string)
				log.Println("[ChatGo]"+UGid.Name+"MessageId:", MessageId)
				err = OpenAI.GetGPTAnswer(StringValue, UGid, MessageId)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}
func PartAdminCommand(Command string) (string, string) {
	Parts := strings.SplitN(Command, "!", 2)
	Parts = strings.SplitN(Parts[1], " ", 2)
	if len(Parts) < 2 {
		return Parts[0], ""
	} else {
		return Parts[0], Parts[1]
	}
}

func PartUserCommand(Command string) (string, string) {
	Parts := strings.SplitN(Command, ".", 2)
	Parts = strings.SplitN(Parts[1], " ", 2)
	if len(Parts) < 2 {
		return Parts[0], ""
	} else {
		return Parts[0], Parts[1]
	}
}

func RunCommand(CommandName, CommandContent, Type string, UGid Type.Id) error {
	if CommandName == "clear" {
		UserCommand.Clear(Type, UGid)
	} else if CommandName == "ChangeModel" {
		UserCommand.ChangeModel(Type, UGid)
	} else if CommandName == "Model" {
		UserCommand.Model(Type, UGid)
	} else if CommandName == "Pre" {
		UserCommand.Premium(Type, UGid)
	} else if CommandName == "help" {
		UserCommand.Help(Type, UGid)
	} else if CommandName == "post" {
		AdminCommand.Post(CommandContent, Type, UGid)
	} else if CommandName == "SetBoard" {
		Sends.SetBoardAllUser("markdown", CommandContent)
	} else if CommandName == "SetPre" {
		AdminCommand.SetUserPremium(CommandContent, UGid)
	} else if CommandName == "AddSu" {
		AdminCommand.AddSu(UGid, CommandContent)
	} else if CommandName == "CutSu" {
		AdminCommand.CutSu(UGid, CommandContent)
	} else if CommandName == "Send" {
		AdminCommand.Send(CommandContent, UGid)
	} else if CommandName == "SetFree" {
		AdminCommand.SetFree(UGid, CommandContent)
	} else if CommandName == "free" {
		UserCommand.CheckFreeTimes(UGid)
	}
	return nil
}

func FindAt(Content string) bool {
	Pattern := `@(Bot|bot|ChatGPTBot|gpt|GPT) [^\s]+`
	Match, _ := regexp.MatchString(Pattern, Content)
	if Match {
		return true
	} else {
		return false
	}
}
