package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	infisical "github.com/infisical/go-sdk"
	"github.com/joho/godotenv"
)

var (
	logs      []string
	logsMutex sync.Mutex
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
	log.Println("Starting server on :8084")
	log.Fatal(http.ListenAndServe("localhost:8084", nil))
}

func addLog(format string, args ...interface{}) {
	logsMutex.Lock()
	defer logsMutex.Unlock()
	logMessage := fmt.Sprintf(format, args...)
	logs = append(logs, logMessage)
}

func addLogWithTimestamp(format string, args ...interface{}) {
	logsMutex.Lock()
	defer logsMutex.Unlock()
	timestamp := time.Now().Format(time.RFC1123)
	logMessage := fmt.Sprintf("[%s] "+format, append([]interface{}{timestamp}, args...)...)
	logs = append(logs, logMessage)
}

func handleLogs(w http.ResponseWriter, r *http.Request) {
	logsMutex.Lock()
	defer logsMutex.Unlock()

	tmpl := template.Must(template.New("logs").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Logs Infisical - Get secrets with Go SDK - Demo Devoxx FR</title>
			<meta http-equiv="refresh" content="1">
		</head>
		<body>
			<h1>Logs Infisical - Get secrets with Go SDK</h1>
			<form action="/clear" method="post">
				<button type="submit">Clear Logs</button>
			</form>
			<pre>{{range .}}{{.}}<br>{{end}}</pre>
		</body>
		</html>
	`))

	tmpl.Execute(w, logs)
}

func handleClearLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		logsMutex.Lock()
		defer logsMutex.Unlock()
		logs = []string{}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
