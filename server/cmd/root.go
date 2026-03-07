package cmd

import (
	"SeeAll/internal/database"
	"SeeAll/internal/metrics"
	"SeeAll/internal/server"
)

func Execute() {

	loadEnv()
	validateEnv()

	runtime := buildRuntime()
	database.InitializeDB()
	metrics.StartFlusher()

	server.Start(runtime)
}
