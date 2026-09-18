package main

import (
	"log"
	"os"

	httpapi "github.com/0xHackerFaith/software-dev-template/go-gin-minimal/internal/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httpapi.NewRouter()
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
