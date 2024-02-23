package greed

import "math/rand"

// A RollFn takes the player's score, the opponent's score, and whether it's the last round, and returns the number of dice to roll.
type RollFn func(self, opponent int, lastRound bool) int

// Possible outcomes of a game.
const (
	player1 = "Player 1"
	player2 = "Player 2"
	Tie     = "Tie"
)

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
func Play(P1, P2 RollFn, loud bool) string {
	var score1, score2 int
	var lastRound bool
	for {
		// TODO: Print the two scores if loud is true.

		p1Dice := P1(score1, score2, lastRound)
		score1 += roll(p1Dice)
		// TODO: Print the number of dice rolled, and the total, if loud is true.

		if score1 > 100 {
			return player2
		}
		if lastRound {
			return determineWinner(score1, score2, loud)
		}
		if p1Dice == 0 {
			lastRound = true
		}

		p2Dice := P2(score2, score1, lastRound)
		score2 += roll(p2Dice)
		// TODO: Print the number of dice rolled, and the total, if loud is true.

		if score2 > 100 {
			return player1
		}

		if lastRound {
			return determineWinner(score1, score2, loud)
		}
		if p2Dice == 0 {
			lastRound = true
		}
	}
}

func determineWinner(score1, score2 int, loud bool) string {
	// TODO: Print the winner if loud is true.
	if score1 > score2 {
		return player1
	} else if score2 > score1 {
		return player2
	} else {
		return Tie
	}
}
