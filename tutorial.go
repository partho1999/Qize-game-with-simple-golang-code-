package main

import "fmt"

func main() {
	fmt.Println("Welcome to the qize game! ")
	fmt.Printf("Enter your Name:")

	var name string
	fmt.Scan(&name)

	fmt.Printf("hello %v, Welcome to the game!\n", name)

	fmt.Printf("Enter your Age:")

	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("yay you can play!!")

	} else {
		fmt.Println("sorry you can't play!!")
		return
	}

	score := 0
	num_qustions := 2

	fmt.Printf("What is better 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+""+answer2 == "RTX 3090" || answer+""+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Wrong!")
	}

	fmt.Printf("how many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("correct!")
		score++
	} else {
		fmt.Println("wrong!")
	}

	fmt.Printf("You scored%v out of %v", score, num_qustions)
	percent := (float64(score) / float64(num_qustions)) * 100
	fmt.Printf("\nYou scored: %v%%.", percent)
}
