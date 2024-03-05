package greed

import (
	"fmt"
	"math/rand"
)

// A RollFn takes the player's score, the opponent's score, and whether it's the last round, and returns the number of dice to roll.
type RollFn func(self, opponent int, lastRound bool) int

// Possible outcomes of a game.
const (
	Player1 = "Player 1"
	Player2 = "Player 2"
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
		if loud {
			fmt.Println("Player 1", score1, "Player 2", score2)
		}
		p1Dice := P1(score1, score2, lastRound)
		score1 += roll(p1Dice)
		// TODO: Print the number of dice rolled, and the total, if loud is true.
		if loud {
			fmt.Println(p1Dice, score1)
		}
		if score1 > 100 {
			Print(loud, "Player 1 Busted")
			return Player2
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
		if loud {
			fmt.Println(p2Dice, score2)
		}
		if score2 > 100 {
			Print(loud, "Player 2 Busted")
			return Player1
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
		return Print(loud, Player1)
	} else if score2 > score1 {
		return Print(loud, Player2)
	} else {
		return Print(loud, Tie)
	}
}

func Print(loud bool, s string) string {
	if loud {
		fmt.Println(s)
	}
	return s
}
