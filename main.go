package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error while reading .env file", err)
	}
	token := os.Getenv("TOKEN")
	fmt.Println(token)
}
