package main

import (
	"log"
	"os"

	"mafiabot/games"

	"github.com/joho/godotenv"
)

func getToken(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func main() {
	games.Token = getToken("token")
	games.Run()
}
