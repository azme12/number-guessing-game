package domain

type Game struct {
	Target       int
	Winner       int
	PlayerTrials map[int]int
	MaxTrials    int
}

func NewGame(maxTrials int) *Game {
	return &Game{
		Target:       1, // default (should be seeded externally)
		PlayerTrials: map[int]int{1: 0, 2: 0},
		MaxTrials:    maxTrials,
	}
}

func (g *Game) Guess(player, value int) (string, bool) {
	if g.Winner != 0 {
		return "Game already won.", true
	}

	g.PlayerTrials[player]++

	if value == g.Target {
		g.Winner = player
		return " Correct guess!", true
	}

	if g.PlayerTrials[1] >= g.MaxTrials && g.PlayerTrials[2] >= g.MaxTrials {
		return " Game over. Both players exhausted trials.", true
	}

	if value < g.Target {
		return "Too low!", false
	}
	return "Too high!", false
}
