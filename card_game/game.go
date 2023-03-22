package main

import (
	"fmt"
	"sort"
)

func isHandOfStraights(hand []int, k int) bool {
	if len(hand)%k != 0 {
		return false
	}

	count := make(map[int]int)
	for _, i := range hand {
		count[i]++
	}

	sort.Ints(hand)
	i := 0
	n := len(hand)

	for i < n {
		current := hand[i]
		for j := 0; j < k; j++ {
			if _, ok := count[current+j]; !ok || count[current+j] == 0 {
				return false
			}
			count[current+j]--
		}
		for i < n && count[hand[i]] == 0 {
			i++
		}
	}
	return true
}

func main() {
	hand := []int{5, 2, 4, 4, 1, 3, 5, 6, 3}
	k := 3
	fmt.Println(isHandOfStraights(hand, k))

	hand2 := []int{1, 9, 3, 5, 7, 4, 2, 9, 11}
	k = 2
	fmt.Println(isHandOfStraights(hand2, k))
}
