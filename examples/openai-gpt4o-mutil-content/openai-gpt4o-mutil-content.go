package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Ideful/langchaingo/llms"
	"github.com/Ideful/langchaingo/llms/openai"
)

func main() {
	llm, err := openai.New(openai.WithModel("gpt-4o"))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are a ocr assistant."),
		llms.TextParts(llms.ChatMessageTypeHuman, "which image is this?"),
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.ImageURLPart("https://github.com/Ideful/langchaingo/blob/main/docs/static/img/parrot-icon.png?raw=true")},
		},
	}

	completion, err := llm.GenerateContent(ctx, content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}
	_ = completion
}
