package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"github.com/xhyrom/swish/pkg/database"
)

func main() {
	fmt.Println("Hello, World!")

	creds := database.GetCredentials()
	fmt.Printf("Credentials are: %+v\n", creds)

	db := database.NewDatabase(creds)
	queue := db.GetQueue()

	fmt.Printf("Queue is: %+v\n", queue)
}