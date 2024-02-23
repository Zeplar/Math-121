package greed

import "fmt"

// inputCheckErrors parses an integer input, and prompts the user to try again if the input is invalid.
func inputCheckErrors() int {
	var nonnegativeNumber int
	_, err := fmt.Scanln(&nonnegativeNumber)
	if err != nil || nonnegativeNumber < 0 {
		fmt.Print("Invalid input. Please enter a nonnegative number: ")

		if err != nil {
			// If the input was gibberish we need to clear the buffer. If it was just negative, we don't.
			var discard string
			fmt.Scanln(&discard)
		}
		// Try again
		return inputCheckErrors()
	}
	return nonnegativeNumber
}

// HumanPlayer prompts the user for input to determine how many dice to roll.
func HumanPlayer(name string) RollFn {
	return func(self, opponent int, lastRound bool) (dice int) {
		if lastRound {
			fmt.Printf("%s, the score is %d to %d and it's the last round. How many dice will you roll? ", name, self, opponent)
		} else {
			fmt.Printf("%s, the score is %d to %d. How many dice will you roll? ", name, self, opponent)
		}
		return inputCheckErrors()
	}
}

// SlowPlayer rolls only one die each turn.
func SlowPlayer(int, int, bool) int {
	return 1
}

// SafePlayer rolls the maximum number of dice that can be safely rolled without busting (going over 100).
func SafePlayer(self, opponent int, lastRound bool) int {
	// TODO: Implement SafePlayer
	return 0
}

// TODO: Implement a new player that you believe can beat the SafePlayer most of the time.
