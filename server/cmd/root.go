package cmd

import (
	"SeeAll/internal/server"
	"os"
)

func Execute() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Start(":" + port)
}
