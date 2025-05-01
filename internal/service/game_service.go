package service

import (
	"math/rand"
	"sync"
	"time"

	"guessinggame/internal/domain"
)

// GameService coordinates the game logic and protects access with a mutex.
type GameService struct {
	Game *domain.Game
	mu   sync.Mutex
}

// NewGameService initializes a new game with a random target number.
func NewGameService() *GameService {
	rand.Seed(time.Now().UnixNano())

	game := domain.NewGame(5) // 5 attempts per player
	game.Target = rand.Intn(100) + 1

	return &GameService{
		Game: game,
	}
}

// Guess processes a guess from the given player and returns a result message and win status.
func (s *GameService) Guess(player int, guess int) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Game.Guess(player, guess)
}

// Reset starts a new game round by resetting the game state and generating a new target number.
func (s *GameService) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Game = domain.NewGame(5)
	s.Game.Target = rand.Intn(100) + 1
}
