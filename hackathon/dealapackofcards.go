package main // piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	x := 0
	for i := 1; i <= 4; i++ {
		fmt.Printf("Player %v: %v, %v, %v\n", i, deck[x], deck[x+1], deck[x+2])
		x += 3
	}
}
