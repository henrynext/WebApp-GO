package main

import (
	"log"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[0:9])

	words := []string{"henry", "nike", "max", "john"}

	log.Println(words)
	log.Println(words[0:2])

}
