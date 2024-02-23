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

// NewHumanPlayer rolls dice according to user input.
func NewHumanPlayer(name string) Player {
	return Player{
		Name: name,
		Roll: inputDice(name),
	}
}

// NewSafePlayer rolls the maximum number of dice that can be safely rolled without busting.
func NewSafePlayer(name string) Player {
	return Player{
		Name: "Safe Player " + name,
		Roll: func(self, _ int, _ bool) int {
			return (100 - self) / 6
		},
	}
}
