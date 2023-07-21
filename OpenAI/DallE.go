package OpenAI

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Type"
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

func CreateImage(UGid Type.Id, Prompt string) (string, error) {
	ApiKey, err := SQLite.GetUserApiKey(UGid.User)
	var Client *openai.Client
	config := openai.DefaultConfig("")
	if ApiKey == "DefaultApiKey" {
		config = openai.DefaultConfig(os.Getenv("DEFAULT_API"))
	} else {
		config = openai.DefaultConfig(ApiKey)
	}
	//ProxyUrl, err := url.Parse(os.Getenv("PROXY"))
	//if err != nil {
	//	return "", err
	//}
	//transport := &http.Transport{
	//	Proxy: http.ProxyURL(ProxyUrl),
	//}
	//config.HTTPClient = &http.Client{
	//	Transport: transport,
	//}

	config.BaseURL = Type.Base
	Client = openai.NewClientWithConfig(config)
	ctx := context.Background()

	ReqUrl := openai.ImageRequest{
		Prompt:         Prompt,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	Resp, err := Client.CreateImage(ctx, ReqUrl)
	if err != nil {
		log.Printf("Image creation error: %v\n\n", err)
		return "", err
	}
	return Resp.Data[0].URL, nil
}
