package main

import (
	"errors"
	"log"
	"os"

	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"github.com/promo/server"
)

var port string

func init() {
	logger, _ := thoth.Init("log")
	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}
	var ok bool
	port, ok = os.LookupEnv("PORT")
	if !ok {
		logger.Log(errors.New("PORT is not set in the file"))
		log.Fatalln("PORT is not set in the file")
	}
}

func main() {

	server.Launch(port)
}
