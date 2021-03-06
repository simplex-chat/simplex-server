package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/simplex-chat/simplex-server/api"
	"github.com/simplex-chat/simplex-server/db"
)

func getPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	db.Open()
	defer db.Close()
	addr, err := getPort()
	if err != nil {
		log.Fatal(err)
	}
	recipientPath := "/connection"
	senderPath := "/connection"
	router := api.New(recipientPath, senderPath)
	log.Printf("Listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
