package main

import "fmt"

type myShelf []int

func main() {
	t := newShelf()

	for _, num := range t {

		if num%2 == 0 {
			fmt.Printf("%d is even \n", num)
		} else {
			fmt.Printf("%d is odd \n", num)
		}

	}
}

func newShelf() myShelf {
	testShelf := myShelf{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	return testShelf
}
