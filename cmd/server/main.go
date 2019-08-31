package main

import (
	"fmt"
	"github.com/iamdejan/ghost-racer-game-server/pkg/httphandler"
	"github.com/iamdejan/ghost-racer-game-server/pkg/modelimpl"
	"log"
	"net/http"
	"os"
	"time"
)

func findListenAddress() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return fmt.Sprintf(":%s", envPort)
	}
	return ":8080"
}

func main() {
	handler := httphandler.Handler{
		RequestMaxDuration: 30 * time.Second,
		Game:               modelimpl.NewGame(),
	}

	listenAddress := findListenAddress()

	log.Fatal(http.ListenAndServe(listenAddress, handler.Routes()))
}