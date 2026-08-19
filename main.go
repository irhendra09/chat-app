package main

import (
	"fmt"
	"log"

	"github.com/irhendra09/chat-app/bootstrap"
	"github.com/irhendra09/chat-app/pkg/env"
)

func main() {
	app := bootstrap.NewAplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "8080"))))
}
