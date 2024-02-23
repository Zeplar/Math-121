package greed

import "fmt"

func inputCheckErrors() int {
	var i int
	_, err := fmt.Scanln(&i)
	if err != nil {
		fmt.Print("Invalid input. Please enter a number: ")
		// Clear the buffer
		var discard string
		fmt.Scanln(&discard)

		// Try again
		return inputCheckErrors()
	}
	return i
}

func inputDice(name string) RollFn {
	return func(self, opponent int, lastRound bool) (dice int) {
		if lastRound {
			fmt.Printf("%s, the score is %d to %d and it's the last round. How many dice will you roll? ", name, self, opponent)
		} else {
			fmt.Printf("%s, the score is %d to %d. How many dice will you roll? ", name, self, opponent)
		}
		return inputCheckErrors()
	}
}

// HumanPlayer rolls dice according to user input.
func NewHumanPlayer(name string) Player {
	return Player{
		Name: name,
		Roll: inputDice(name),
	}
}

// SlowPlayer rolls only one die each turn.
func NewSlowPlayer(name string) Player {
	return Player{
		Name: name,
		Roll: func(int, int, bool) int {
			return 1
		},
	}
}

// SafePlayer rolls the maximum number of dice that can be safely rolled without busting (going over 100).
func NewSafePlayer(name string) Player {
	return Player{
		Name: "Safe Player " + name,
		// TODO: Implement the Roll function for the SafePlayer.
	}
}

// TODO: Implement a new player that you believe can beat the SafePlayer most of the time.
