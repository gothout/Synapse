package env

import (
	"log"
	"os"
)

func GetPortServer() string {
	if os.Getenv("PORT") == "" {
		log.Fatal("Env PORT nao setado")
		os.Exit(1)
	}
	return os.Getenv("PORT")
}

func GetHostServer() string {
	if os.Getenv("HOST") == "" {
		log.Fatal("Env HOST nao setado")
		os.Exit(1)
	}
	return os.Getenv("HOST")
}

func GetSecretKey() string {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("SECRET_KEY not set")
	}
	return secret
}
