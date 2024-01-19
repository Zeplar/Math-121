package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		fmt.Println("zzz", i)
		time.Sleep(time.Second)
		i++
		if i >= 5 {
			break
		}
	}

	i = 0
	for i < 5 {
		fmt.Println("zzz", i)
		time.Sleep(time.Second)
		i++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("zzz", i)
		time.Sleep(time.Second)
	}

}
