package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LockBlock-dev/javascriptdle/controller"
	"github.com/LockBlock-dev/javascriptdle/repository"
	"github.com/LockBlock-dev/javascriptdle/utils"
)

const HTTP_PORT = 8080

func main() {
	log.Println("--- JavaScriptDle ---")

	guessableRepo := repository.NewMemoryGuessableRepository()
	guessController := controller.NewGuessController(guessableRepo)

	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		utils.GenerateTodaysGuess(guessableRepo)

		for range ticker.C {
			utils.GenerateTodaysGuess(guessableRepo)
		}
	}()

	http.HandleFunc("GET /", guessController.Home)
	http.HandleFunc("POST /guess", guessController.Guess)

	http.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Printf("Started listening on http://localhost:%d\n", HTTP_PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", HTTP_PORT), nil))
}
