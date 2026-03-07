package cmd

import (
	"SeeAll/internal/server"
)

func Execute() {

	loadEnv()
	validateEnv()

	runtime := buildRuntime()

	server.Start(runtime)
}
