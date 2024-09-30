package infra

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func Run(port string, geminiKey string) error {
	app := fiber.New(fiber.Config{
		AppName:      "Socratic code",
		ServerHeader: "Fiber",
	})

	gClient, err := genai.NewClient(context.Background(), option.WithAPIKey(geminiKey))
	if err != nil {
		return err
	}
	textModel := gClient.GenerativeModel("gemini-1.5-flash")
	textModel.ResponseMIMEType = "application/json"
	textModel.ResponseSchema = &genai.Schema{
		Type: genai.TypeArray,
		Items: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"line": {Type: genai.TypeInteger},
				"q":    {Type: genai.TypeString},
			},
		},
	}

	server := NewServer(app, textModel)

	err = server.App.Listen(fmt.Sprintf(":%v", port))
	return err
}
