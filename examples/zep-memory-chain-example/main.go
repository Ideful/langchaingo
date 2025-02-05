package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Ideful/langchaingo/chains"
	"github.com/Ideful/langchaingo/llms/openai"
	zepLangchainMemory "github.com/Ideful/langchaingo/memory/zep"
	"github.com/getzep/zep-go"
	zepClient "github.com/getzep/zep-go/client"
	zepOption "github.com/getzep/zep-go/option"
)

func main() {
	ctx := context.Background()

	client := zepClient.NewClient(zepOption.WithAPIKey(os.Getenv("ZEP_API_KEY")))
	sessionID := os.Getenv("ZEP_SESSION_ID")

	llm, err := openai.New()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	c := chains.NewConversation(
		llm,
		zepLangchainMemory.NewMemory(
			client,
			sessionID,
			zepLangchainMemory.WithMemoryType(zep.MemoryGetRequestMemoryTypePerpetual),
			zepLangchainMemory.WithReturnMessages(true),
			zepLangchainMemory.WithAIPrefix("Robot"),
			zepLangchainMemory.WithHumanPrefix("Joe"),
		),
	)
	res, err := chains.Run(ctx, c, "Hi! I'm John Doe")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("Response: %s\n", res)

	res, err = chains.Run(ctx, c, "What is my name?")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("Response: %s\n", res)
}
