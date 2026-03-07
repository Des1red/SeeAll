package cmd

import (
	"log"
	"os"
	"strings"
)

func loadEnv() {
	data, err := os.ReadFile(".env")
	if err != nil {
		log.Println("no .env file found")
		return
	}

	for _, line := range strings.Split(string(data), "\n") {

		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// don't override existing env vars
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}
}
