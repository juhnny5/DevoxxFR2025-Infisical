package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	infisical "github.com/infisical/go-sdk"
	"github.com/joho/godotenv"
)

var (
	logs        []string
	logsMutex   sync.Mutex
	stopRefresh = false
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

	go func() {
		for {
			if stopRefresh {
				break
			}

			apiKeySecret, err := client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
				SecretKey:   "API_KEY",
				Environment: "dev",
				ProjectID:   projectID,
				SecretPath:  "/",
			})

			if err != nil {
				addLog("Error retrieving API Key Secret: %v", err)
			} else {
				addLogWithTimestamp("(only) API Key Secret: %v", apiKeySecret.SecretValue)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	http.HandleFunc("/", handleLogs)
	http.HandleFunc("/logs", handleLogs)
	http.HandleFunc("/clear", handleClearLogs)
	http.HandleFunc("/stop", handleStopRefresh)
	log.Println("Starting server on http://localhost:8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
