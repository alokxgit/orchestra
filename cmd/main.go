package main

import (

	"github.com/alokxcode/orchestra/agent"
)

func main() {
	
	Agent := agent.New(agent.Gemini,"someapikey")
	Agent.Invoke()
}

