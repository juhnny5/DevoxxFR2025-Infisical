package main

import (
	"context"
	"log"
	"os"
	"time"

	infisical "github.com/infisical/go-sdk"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	clientID := os.Getenv("CLIENT_ID")
	siteURL := os.Getenv("SITE_URL")
	clientSecret := os.Getenv("CLIENT_SECRET")
	projectID := os.Getenv("PROJECT_ID")

	client := infisical.NewInfisicalClient(context.Background(), infisical.Config{
		SiteUrl:          siteURL,
		AutoTokenRefresh: true,
	})

	_, err = client.Auth().UniversalAuthLogin(clientID, clientSecret)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	for {
		apiKeySecret, err := client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
			SecretKey:   "API_KEY",
			Environment: "dev",
			ProjectID:   projectID,
			SecretPath:  "/",
		})

		if err != nil {
			log.Printf("Error retrieving API Key Secret: %v", err)
		} else {
			log.Printf("API Key Secret: %v", apiKeySecret)
		}

		time.Sleep(2 * time.Second)
	}
}
