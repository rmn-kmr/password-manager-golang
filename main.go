package main

import (
	"log"
	"passwordmanager/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error occured while loading env.\nErr: %s", err)
	}
	err = server.Server()
	if err != nil {
		log.Print(err)
	}
}
