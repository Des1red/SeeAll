package cmd

import (
	"SeeAll/internal/model"
	"log"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func validateEnv() {

	dev := os.Getenv("DEV")
	apiDev := os.Getenv("API_DEV")
	apiProd := os.Getenv("API_PROD")
	devOrigin := os.Getenv("DEV_ORIGIN")
	prodOrigin := os.Getenv("PROD_ORIGIN")
	user := os.Getenv("ADMIN_USER")
	pass := os.Getenv("ADMIN_PASS")
	secret := os.Getenv("JWT_SECRET")
	hashsalt := os.Getenv("HASH_SALT")

	if hashsalt == "" {
		log.Fatal("ENV HASH_SALT is missing")
	}

	if secret == "" {
		log.Fatal("ENV JWT_SECRET is missing")
	}
	if user == "" || pass == "" {
		log.Fatal("ENV USER or ENV PASS is missing")
	}

	if dev != "true" && dev != "false" {
		log.Fatal("ENV DEV must be 'true' or 'false'")
	}

	if dev == "true" && apiDev == "" {
		log.Fatal("ENV API_DEV is not set")
	}

	if apiProd == "" {
		log.Fatal("ENV API_PROD is not set")
	}

	if dev == "true" && devOrigin == "" {
		log.Fatal("ENV DEV_ORIGIN is not set")
	}

	if prodOrigin == "" {
		log.Fatal("ENV PROD_ORIGIN is not set")
	}
}

func buildRuntime() model.Runtime {

	return model.Runtime{
		Dev:        os.Getenv("DEV") == "true",
		APIDev:     os.Getenv("API_DEV"),
		APIProd:    os.Getenv("API_PROD"),
		DevOrigin:  os.Getenv("DEV_ORIGIN"),
		ProdOrigin: os.Getenv("PROD_ORIGIN"),
		Port:       ":" + getPort(),
		User:       os.Getenv("ADMIN_USER"),
		Pass:       os.Getenv("ADMIN_PASS"),
		JWTsecret:  os.Getenv("JWT_SECRET"),
		HashSalt:   os.Getenv("HASH_SALT"),
	}
}
