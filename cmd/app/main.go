package main

import (
	"log"
	"os"

	"github.com/ary82/socratic-leetcode/internal/infra"
	"github.com/joho/godotenv"
)

func main() {
	mode := os.Getenv("MODE")
	if mode != "prod" {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatal(err)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port envvar empty")
	}
	geminiKey := os.Getenv("GEMINI_KEY")
	if geminiKey == "" {
		log.Fatal("gemini_key envvar empty")
	}

	err := infra.Run(port, geminiKey)
	if err != nil {
		log.Fatal(err)
	}
}
