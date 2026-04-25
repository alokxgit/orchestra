package main

import (
	"context"
	"fmt"

	"github.com/alokxcode/orchestra/agent"
)

func main() {
	
	Agent := agent.New("anthropic","someapikey")
	cfg := agent.ChatConfig {
		Model: "anthropic",
		MaxToken: 200,
		SystemPrompt: "Hello",
	}
	prmpt := agent.Prompt {
		Text : "Hii",
	}
	res,err := Agent.Invoke(&prmpt, "some memory", &cfg , context.Background()  )
	if err != nil {
		return
	}
	fmt.Println(res)
}

