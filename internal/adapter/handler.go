package adapter

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"guessinggame/internal/service"
)

type Handler struct {
	GameSvc *service.GameService
	Tmpl    *template.Template
}

// NewHandler initializes a new HTTP handler with the game service and parsed HTML template.
func NewHandler(svc *service.GameService) *Handler {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	return &Handler{
		GameSvc: svc,
		Tmpl:    tmpl,
	}
}

// Home handles the root route and renders the game UI.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}
}

// Guess handles player guesses via query parameters and returns a JSON response.
func (h *Handler) Guess(w http.ResponseWriter, r *http.Request) {
	playerID, err1 := strconv.Atoi(r.URL.Query().Get("player"))
	guessVal, err2 := strconv.Atoi(r.URL.Query().Get("guess"))

	if err1 != nil || err2 != nil || (playerID != 1 && playerID != 2) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	message, won := h.GameSvc.Guess(playerID, guessVal)

	if won {
		go func() {
			time.Sleep(3 * time.Second)
			h.GameSvc.Reset()
		}()
	}

	response := map[string]interface{}{
		"message": message,
		"winner":  won,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
