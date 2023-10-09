//config.go
// author: rjjm94 on 10/09/2023
// functions: SetupEnvironment
// variables: redisAddr, redisPassword, redisDB
// struct: Settings
// description: This file contains the functionality to setup the environment.

package config

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"github.com/go-redis/redis"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Settings represents the settings from the settings sheet.
type Settings struct {
	PromptSystemMessage            string
	PromptTemperature              float32
	PromptMaxTokens                int
	SheetName                      string
	MessageVariable1               string
	GenerateMessageVariablesColumn string
	PromptColumn                   string
	PromptResponseColumn           string
	Variable1Column                string
	Variable1PromptSystemMessage   string
	Variable1PlacedColumn          string
	Variable1MaxTokens             int
	Variable1Temperature           float32
}

// SetupEnvironment loads environment variables, creates the Google Sheets and Redis services, and returns the context, Redis client and Sheets service.
// It decodes the Google service account key, parses it and creates a new Google Sheets service.
// It also creates a new Redis client and clears the Redis database.
func SetupEnvironment() (context.Context, *redis.Client, *sheets.Service) {
	ctx := context.Background()

	// Decode the Google service account key
	b, err := base64.StdEncoding.DecodeString(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		log.Fatalf("Error decoding service account key: %v", err)
	}

	// Parse the Google service account key
	conf, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Error parsing service account key: %v", err)
	}

	// Create a new Google Sheets service
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(conf.Client(ctx)))
	if err != nil {
		log.Fatalf("Error creating new Sheets service: %v", err)
	}

	// Create a new Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Clear the Redis database
	err = redisClient.FlushDB().Err()
	if err != nil {
		log.Fatalf("Error clearing Redis database: %v", err)
	}

	return ctx, redisClient, srv
}
