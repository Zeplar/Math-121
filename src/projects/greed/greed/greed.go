package greed

import "math/rand"

type RollFn func(self, opponent int, lastRound bool) int

type Player struct {
	Name  string
	Score int
	// Roll takes the players' current scores and returns the number of dice to roll. lastRound is true if the opponent has passed.
	Roll RollFn
}

// roll rolls the given number of D6 and returns the total.
func roll(dice int) int {
	total := 0
	for i := 0; i < dice; i++ {
		total += rand.Intn(6) + 1
	}
	return total
}

// Play runs the game until a player wins, and outputs the winner's name.
// If loud is true, it prints the game state after each turn.
func Play(P1, P2 Player, loud bool) string {
	lastRound := false
	for {
		// TODO: Print the two scores if loud is true.

		p1Dice := P1.Roll(P1.Score, P2.Score, lastRound)
		P1.Score += roll(p1Dice)
		// TODO: Print the number of dice rolled, and the total, if loud is true.

		if P1.Score > 100 {
			return P2.Name
		}
		if lastRound {
			return scoreGame(P1, P2, loud)
		}
		if p1Dice == 0 {
			lastRound = true
		}

		p2Dice := P2.Roll(P2.Score, P1.Score, lastRound)
		P2.Score += roll(p2Dice)
		// TODO: Print the number of dice rolled, and the total, if loud is true.

		if P2.Score > 100 {
			return P1.Name
		}

		if lastRound {
			return scoreGame(P1, P2, loud)
		}
		if p2Dice == 0 {
			lastRound = true
		}
	}
}

// scoreGame ouputs the winner of the game, or "Tie".
func scoreGame(P1, P2 Player, loud bool) string {
	if P1.Score > P2.Score {
		return P1.Name
	} else if P2.Score > P1.Score {
		return P2.Name
	} else {
		return "Tie"
	}
}
