package devmode

import (
	"SeeAll/internal/model"
	"fmt"
)

func InitDev(r model.Runtime) {
	setDev(r.Dev)
	if !isDev() {
		return
	}
	fmt.Println("Runtime configuration:")
	fmt.Println("----------------------")

	fmt.Println("Dev:", r.Dev)
	fmt.Println("API_DEV:", r.APIDev)
	fmt.Println("API_PROD:", r.APIProd)
	fmt.Println("DEV_ORIGIN:", r.DevOrigin)
	fmt.Println("PROD_ORIGIN:", r.ProdOrigin)
	fmt.Println("PORT:", r.Port)
	fmt.Println("ADMIN_USER:", r.User)
	fmt.Println("ADMIN_PASS:", r.Pass)
	fmt.Println("JWT_SECRET:", r.JWTsecret)
	fmt.Println("HASH_SALT:", r.HashSalt)
}
