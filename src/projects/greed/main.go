package main

import (
	"fmt"

	"github.com/Zeplar/Math-121/src/projects/greed/greed"
)

func main() {
	p1 := greed.SafePlayer
	p2 := greed.Risky
	win1 := 0
	win2 := 0
	tie := 0
	for i := 0; i < 10000; i++ {
		winner := greed.Play(p1, p2, false)
		if winner == greed.Player1 {
			win1++
		} else if winner == greed.Player2 {
			win2++
		} else if winner == greed.Tie {
			tie++
		}
	}
	fmt.Println(win1, win2, tie)
}

/*
	------- HOMEWORK ---------

	Inside greed/greed.go, implement the TODO lines regarding the `loud` feature in the Play function. Test it by running a game between two SlowPlayers.

	Inside greed/players.go, implement the SafePlayer.

	Inside greed/players.go, implement a new player that you believe can beat the SafePlayer most of the time.

	Determine whether there is a first-player advantage in a game between two SafePlayers.

*/
