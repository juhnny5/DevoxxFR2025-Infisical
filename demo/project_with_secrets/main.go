package main

import (
	"fmt"
	"os"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	fmt.Printf("Connecting to the database with user: %s\n", dbUser)

	if dbUser != "" && dbPassword != "" {
		fmt.Println("Connected successfully!")
	} else {
		fmt.Println("Failed to connect: missing credentials")
	}
}
