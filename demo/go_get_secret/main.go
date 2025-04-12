package main

import (
	"context"
	"fmt"
	"os"

	infisical "github.com/infisical/go-sdk"
)

func main() {

	client := infisical.NewInfisicalClient(context.Background(), infisical.Config{
		SiteUrl:          "https://infisical.devoxx.jbriault.fr",
		AutoTokenRefresh: true,
	})

	_, err := client.Auth().UniversalAuthLogin("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")

	if err != nil {
		fmt.Printf("Authentication failed: %v", err)
		os.Exit(1)
	}

	apiKeySecret, err := client.Secrets().Retrieve(infisical.RetrieveSecretOptions{
		SecretKey:   "API_KEY",
		Environment: "dev",
		ProjectID:   "YOUR_PROJECT_ID",
		SecretPath:  "/",
	})

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("API Key Secret: %v", apiKeySecret)

}
