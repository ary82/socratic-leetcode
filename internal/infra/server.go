package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/generative-ai-go/genai"
)

type Server struct {
	App       *fiber.App
	TextModel *genai.GenerativeModel
}

func NewServer(app *fiber.App, textModel *genai.GenerativeModel) *Server {
	s := &Server{
		App:       app,
		TextModel: textModel,
	}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	s.App.Use(logger.New())
	s.App.Use(compress.New())

	s.App.Post("/generate", s.getSocraticQuestions)
}

func (s *Server) getSocraticQuestions(ctx *fiber.Ctx) error {
	req := new(TextGenerationRequest)

	err := ctx.BodyParser(req)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}
	err = req.Validate()
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	dynamicPrompt := fmt.Sprintf(PROMPT, req.Question, req.Lang, req.Loc, req.Code)
	res, err := s.TextModel.GenerateContent(
		context.Background(),
		genai.Text(dynamicPrompt),
	)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	resJson := []TextGenerationResponse{}
	err = json.Unmarshal([]byte(res.Candidates[0].Content.Parts[0].(genai.Text)), &resJson)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(resJson)
}
