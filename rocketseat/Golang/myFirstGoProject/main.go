package main

import "fmt"

func main() {
	fmt.Println("Jogo da Adivinhação")
}

func helper(guess, number int) []int {

	numbers := []int{}

	if guess < number {
		for i := guess; i <= number; i++ {
			numbers = append(numbers, i)

		}
	} else {
		for i := guess; i >= number; i-- {
			numbers = append(numbers, i)

		}

	}
	return numbers
}
