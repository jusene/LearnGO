package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() (left int) {
	var use int
	var dis map[string]int
	for _, user := range users {
		dis = getCoins(user)
	}

	for _, coins := range dis {
		use += coins
	}

	return coins - use
}

func getCoins(user string) map[string]int {
	m := map[string]int{
		"e":1,
		"i":2,
		"o":3,
		"u":4,
	}

	var sum int

	a := func(s string) int {
		return strings.Count(strings.ToLower(user), s)
	}
	for k := range m {
		sum += a(k) * m[k]
	}
	distribution[user] = sum
	return distribution
}