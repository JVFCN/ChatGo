package OpenAI

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"context"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/yhchat/bot-go-sdk/openapi"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var RecentDataList1 = RecentDataList{}

func GetGPTAnswer(Prompt string, UGid Type.Id, MsgId string) error {
	err := godotenv.Load("data/.env")
	if err != nil {
		return err
	}
	err = SQLite.AddUser(UGid.User)
	log.Println("[ChatGo]尝试添加用户:" + UGid.User)
	Model := SQLite.GetUserModel(UGid.User)
	if err != nil {
		return err
	}
	if Model == "gpt-4" || Model == "gpt-4-32k" || Model == "gpt-3.5-turbo" || Model == "gpt-3.5-turbo-16k" {
		log.Println("[ChatGo]" + UGid.Name + "用的是付费模型")
		if SQLite.GetUserFreeTimes(UGid.User) <= 0 {
			log.Println("[ChatGo]" + UGid.Name + "免费次数不足")
			if SQLite.IsPremium(UGid.User) == false {
				Buttons := []openapi.Button{
					{
						Text:       "购买会员",
						ActionType: 3,
						Value:      "buy" + UGid.User + "|" + UGid.MainType,
					},
				}
				_, err = Sends.EditTextMessage(MsgId, UGid.MainId, UGid.MainType, "您不是高级用户, 无法使用该模型", Buttons)
				return nil
			} else {
				if SQLite.GetUserPremiumExpire(UGid.User) <= time.Now().Unix() {
					log.Println("[ChatGo]" + UGid.Name + "会员过期")
					Buttons := []openapi.Button{
						{
							Text:       "续费会员",
							ActionType: 3,
							Value:      "buy" + UGid.User + "|" + UGid.MainType,
						},
					}
					_, err = Sends.EditTextMessage(MsgId, UGid.MainId, UGid.MainType, "您的会员已过期, 无法使用该模型", Buttons)
					return nil
				} else {
					log.Println("[ChatGo]" + UGid.Name + "是会员")
					err := answer(Prompt, UGid, MsgId)
					if err != nil {
						return err
					}
				}
			}
		}
		err := SQLite.UpdateUserFreeTimes(UGid.User, SQLite.GetUserFreeTimes(UGid.User)-1)
		log.Println("[ChatGo]" + UGid.Name + "免费次数减一, 还剩" + strconv.Itoa(SQLite.GetUserFreeTimes(UGid.User)) + "次")
		if err != nil {
			return err
		}
	}
	log.Println("[ChatGo]" + UGid.Name + "没有会员")
	err = answer(Prompt, UGid, MsgId)
	if err != nil {
		return err
	}
	return nil
}

func answer(Prompt string, UGid Type.Id, MsgId string) error {
	ApiKey, err := SQLite.GetUserApiKey(UGid.User)
	var Client *openai.Client
	config := openai.DefaultConfig("")
	if ApiKey == "DefaultApiKey" {
		config = openai.DefaultConfig(os.Getenv("DEFAULT_API"))
	} else {
		config = openai.DefaultConfig(ApiKey)
	}
	ProxyUrl := os.Getenv("PROXY")

	if ProxyUrl != "" {
		ProxyUrlParse, err := url.Parse(ProxyUrl)
		if err != nil {
			log.Println(err)
			return err
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(ProxyUrlParse),
		}
		config.HTTPClient = &http.Client{
			Transport: transport,
		}
	}

	config.BaseURL = Type.Base
	Client = openai.NewClientWithConfig(config)

	Messages, err := SQLite.GetUserContext(UGid.User)
	if err != nil {
		log.Println(err)
		return err
	}
	var AllMsg []openai.ChatCompletionMessage

	for _, index := range Messages {
		AllMsg = append(AllMsg, openai.ChatCompletionMessage{
			Role:    index.Role,
			Content: index.Content,
		})
	}

	AllMsg = append(AllMsg, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: Prompt,
	})

	Resp, err := Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: AllMsg,
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return err
	}
	AnswerContent := Resp.Choices[0].Message.Content

	log.Println("[ChatGo]回答(User:" + UGid.Name + "): " + AnswerContent)

	Buttons := []openapi.Button{
		{
			Text:       "复制回答",
			ActionType: 2,
			Value:      AnswerContent,
		},
		{
			Text:       "翻译",
			ActionType: 3,
			Value:      "translate" + AnswerContent,
		},
		{
			Text:       "重新响应",
			ActionType: 3,
			Value:      "AgainReply" + Prompt,
		},
	}

	RecentDataList1.AddDataItem(DataItem{Timestamp: time.Now(), Value: UGid.Name + " 问:" + Prompt})

	_, err = Sends.EditTextMessage(MsgId, UGid.MainId, UGid.MainType, AnswerContent, Buttons)
	if err != nil {
		return err
	}

	jsonString, err := json.Marshal(Messages)
	if err != nil {
		return err
	}

	var NewData []Type.Data
	err = json.Unmarshal(jsonString, &NewData)
	if err != nil {
		return err
	}
	NewData = append(NewData, Type.Data{
		Role:    openai.ChatMessageRoleUser,
		Content: Prompt,
	})
	NewData = append(NewData, Type.Data{
		Role:    openai.ChatMessageRoleAssistant,
		Content: AnswerContent,
	})

	NewJsonString, err := json.Marshal(NewData)
	if err != nil {
		return err
	}

	err = SQLite.UpdateUserContext(UGid.User, string(NewJsonString))
	if err != nil {
		return err
	}
	return nil
}

func (rdl *RecentDataList) AddDataItem(item DataItem) {
	rdl.data = append(rdl.data, item)
}

func (rdl *RecentDataList) GetRecentData() []DataItem {
	var recentData []DataItem
	oneHourAgo := time.Now().Add(-time.Hour)

	for _, item := range rdl.data {
		if item.Timestamp.After(oneHourAgo) {
			recentData = append(recentData, item)
		}
	}

	return recentData
}

type DataItem struct {
	Timestamp time.Time
	Value     string
}

type RecentDataList struct {
	data []DataItem
}
