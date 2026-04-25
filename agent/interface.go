package agent

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"google.golang.org/genai"
)

type ChatConfig struct {
	Model string
	SystemPrompt string
	MaxToken int
}


type Prompt struct {
	Text string
	Image string
	Audio string
	Video string
	Clipboard string
}


type Agent interface {
	Invoke(prompt *Prompt,Memory any, cfg *ChatConfig, ctx context )
}
type Client struct {
	Gemini genai.Client
	Anthropic anthropic.Client
}

type Adaptor struct {
	Client Client
}

type Provider string
const (
	Gemini Provider = "gemini"
	Anthropic Provider = "anthropic"
	OpenAi Provider = "openai"
)

func New(provider Provider, api_key string) Agent {
	switch provider {
	case Gemini:
		return newGemini(api_key)
	case Anthropic:
		return newAnthropic(api_key)
	case OpenAi:
		return newOpenai(api_key)
	default:
		return Agent{}
	}
}

func newGemini(api_key string) Agent {

	cc := genai.ClientConfig{
		APIKey: api_key,
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &cc)
	if err != nil {
		fmt.Println("-----------------Error while init-------------", err)
		return &Adaptor{}
	}

	adaptor := Adaptor{
		Client : Client{Gemini: *client},
	}

	return &adaptor
}
