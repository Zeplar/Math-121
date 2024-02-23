package greed

import "math/rand"

type Player struct {
	Name  string
	Score int
	// Roll takes the players' current scores and returns the number of dice to roll.
	Roll func(self, opponent int) int
}

// Roll rolls the given number of D6 and returns the total.
func Roll(dice int) int {
	total := 0
	for i := 0; i < dice; i++ {
		total += rand.Intn(6) + 1
	}
	return total
}

// Play runs the game until a player wins, and outputs the winner's name.
// If loud is true, it prints the game state after each turn.
func Play(P1, P2 Player, loud bool) string {
	for {
		p1Dice := P1.Roll(P1.Score, P2.Score)
		P1.Score += Roll(p1Dice)
		if P1.Score > 100 {
			return P2.Name
		}

		p2Dice := P2.Roll(P2.Score, P1.Score)
		P2.Score += Roll(p2Dice)
		if P2.Score > 100 {
			return P1.Name
		}

		if p1Dice == 0 && p2Dice == 0 {
			if P1.Score > P2.Score {
				return P1.Name
			} else if P2.Score > P1.Score {
				return P2.Name
			} else {
				return "Tie"
			}
		}
	}
}
