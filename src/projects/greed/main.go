package main

import (
	"fmt"

	"github.com/Zeplar/Math-121/src/projects/greed/greed"
)

func main() {
	p1 := greed.HumanPlayer("Mitchell")
	p2 := greed.HumanPlayer("Josh")

	winner := greed.Play(p1, p2, false)
	fmt.Printf("The winner is %s\n", winner)
}

/*
	------- HOMEWORK ---------

	Inside greed/greed.go, implement the TODO lines regarding the `loud` feature in the Play function. Test it by running a game between two SlowPlayers.

	Inside greed/players.go, implement the SafePlayer.

	Inside greed/players.go, implement a new player that you believe can beat the SafePlayer most of the time.

	Determine whether there is a first-player advantage in a game between two SafePlayers.

*/
