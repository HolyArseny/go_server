package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
);

type Env struct {
	HTTP_PORT string
}

func GetEnv () Env {
	err := godotenv.Load();

  if err != nil {
    log.Fatal("Error loading .env file")
  }

	config := Env{ HTTP_PORT: os.Getenv("HTTP_PORT") };

	return config;
}