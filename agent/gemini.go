package agent

import (
	"context"

	"google.golang.org/genai"
)


func (a *Adaptor) Invoke(ctx context.Context, req []Msg, systemPrompt string) (string, error) {
	ctx = context.Background()
	var Content []*genai.Content
	for _, r := range req {
		var parts []*genai.Part
		for _, m := range r.Content {
			msg := genai.Part{Text: m}
			parts = append(parts, &msg)
		}
		cont := genai.Content{Role: r.Role, Parts: parts}
		Content = append(Content, &cont)
	}

	modelConfig := genai.GenerateContentConfig{MaxOutputTokens: 100}

	res, err := a.Client.Gemini.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash-lite",
		Content,
		&modelConfig,
	)
	if err != nil {
		return "", err
	}
	return res.Text(), nil
}
