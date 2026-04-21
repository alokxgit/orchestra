package agent

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
)


func (a *Adaptor) Invoke(ctx context.Context, req []Msg, systemPrompt string) {

	var messages []anthropic.MessageParam
	// loops over each msgs - [ [ user : "hii" ], [ ai : "hello" ] ]
	for _, r := range req {
		var contentBlock []anthropic.ContentBlockParamUnion
		// loops over each different types of msgs - [ " user : 'what's the promlem in this code' ", IMG, CODE_SNIPPET ]
		for _, m := range r.Content {
			contentBlock = append(contentBlock, anthropic.NewTextBlock(m))
		}
		msg := anthropic.MessageParam{
			Content: contentBlock,
			Role:    anthropic.MessageParamRole(r.Role),
		}
		messages = append(messages, msg)

	}

	systemPromptStruct := anthropic.TextBlockParam{
		Text: systemPrompt,
	}

	systemPromptSlice := []anthropic.TextBlockParam{systemPromptStruct}

	message := anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeHaiku4_5,
		Messages:  messages,
		System:    systemPromptSlice,
		MaxTokens: 1000,
	}
	a.Client.Anthropic.Messages.New(ctx, message)
}
