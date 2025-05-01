package main

import (
	"fmt"
	"log"
	"net/http"

	"guessinggame/internal/adapter"
	"guessinggame/internal/service"
)

func main() {
	gameSvc := service.NewGameService()
	handler := adapter.NewHandler(gameSvc)

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/guess", handler.Guess)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
