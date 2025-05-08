package env

import (
	"log"
	"os"
)

func GetDatabaseHost() string {
	if os.Getenv("DATABASE_HOST") == "" {
		log.Fatal("DATABASE_HOST not set")
		os.Exit(1)
	}
	return os.Getenv("DATABASE_HOST")
}

func GetDatabasePort() string {
	if os.Getenv("DATABASE_PORT") == "" {
		log.Fatal("DATABASE_PORT not set")
		os.Exit(1)
	}

	return os.Getenv("DATABASE_PORT")
}

func GetDatabaseName() string {
	if os.Getenv("DATABASE_NAME") == "" {
		log.Fatal("DATABASE_NAME not set")
		os.Exit(1)
	}

	return os.Getenv("DATABASE_NAME")
}

func GetDatabaseUser() string {
	if os.Getenv("DATABASE_USER") == "" {
		log.Fatal("DATABASE_USER not set")
		os.Exit(1)
	}

	return os.Getenv("DATABASE_USER")
}

func GetDatabasePassword() string {
	if os.Getenv("DATABASE_PASSWORD") == "" {
		log.Fatal("DATABASE_PASSWORD not set")
		os.Exit(1)
	}

	return os.Getenv("DATABASE_PASSWORD")
}
