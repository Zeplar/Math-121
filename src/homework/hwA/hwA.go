package main

type Game struct {
	P1 Player
	P2 Player
}

type Player struct {
	Name  string
	Score int
	// Roll takes the players' current scores and returns the number of dice to roll.
	Roll func(self, opponent int) int
}

func main() {
	var g Game
}
