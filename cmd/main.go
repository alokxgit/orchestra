package main

import (
	"context"
	"fmt"

	"github.com/alokxcode/orchestra/agent"
)

func main() {
	
	Agent := agent.New("gemini","AIzaSyAUOBeDoSQyG66gCq7Vfwo5I1rGVwifsKA")
	cfg := agent.ChatConfig {
		Model: "gemini-2.0-flash-lite",
		MaxToken: 200,
		SystemPrompt: "Hello",
	}
	prmpt := agent.Prompt {
		Text : "Hii",
	}
	content := map[string]string{"text":"hello"}
	history := agent.Message{
		Role: "user",
		Content: content,
	}
	memory := []agent.Message{history}
	res,err := Agent.Invoke(prmpt, memory, &cfg , context.Background()  )
	if err != nil {
		fmt.Println("error",err)
		return
	}
	fmt.Println(res)
}

