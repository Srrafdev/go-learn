package main// piscine

type food struct {
	burger  int
	chips   int
	nuggets int
}

func FoodDeliveryTime(order string) int {
	var time food
	time.burger = 15
	time.chips = 10
	time.nuggets = 12
	if order == "burger" {
		return time.burger
	}
	if order == "chips" {
		return time.chips
	}
	if order == "nuggets" {
		return time.nuggets
	}
	return 404
}
