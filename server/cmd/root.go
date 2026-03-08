package cmd

import (
	"SeeAll/internal/database"
	"SeeAll/internal/devmode"
	"SeeAll/internal/metrics"
	"SeeAll/internal/server"
)

func Execute() {

	loadEnv()
	validateEnv()

	runtime := buildRuntime()
	devmode.InitDev(runtime)
	database.InitializeDB()
	metrics.StartFlusher()
	startPrewarm()
	server.Start(runtime)
}
