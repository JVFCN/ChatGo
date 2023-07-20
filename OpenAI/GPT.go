package OpenAI

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/yhchat/bot-go-sdk/openapi"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetGPTAnswer(Prompt string, UGid Type.Id, MsgId string) error {
	err := godotenv.Load("data/.env")
	if err != nil {
		return err
	}
	err = SQLite.AddUser(UGid.User)

	if SQLite.GetUserModel(UGid.User) == "gpt-4" || SQLite.GetUserModel(UGid.User) == "gpt-4-32k" {
		if SQLite.GetUserFreeTimes(UGid.User) <= 0 {
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
					fmt.Println("Premium")
					err := Answer(Prompt, UGid, MsgId)
					if err != nil {
						return err
					}

				}
			}
		}
	} else {
		fmt.Println("NoPremium")
		err := Answer(Prompt, UGid, MsgId)
		if err != nil {
			return err
		}
	}
	return nil
}

func Answer(Prompt string, UGid Type.Id, MsgId string) error {
	ApiKey, err := SQLite.GetUserApiKey(UGid.User)
	var Client *openai.Client
	config := openai.DefaultConfig("")
	if ApiKey == "DefaultApiKey" {
		config = openai.DefaultConfig(os.Getenv("DEFAULT_API"))
	} else {
		config = openai.DefaultConfig(ApiKey)
	}
	ProxyUrl, err := url.Parse(os.Getenv("PROXY"))
	if err != nil {
		return err
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(ProxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}

	config.BaseURL = "https://api.mctools.online/v1"
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
