package main

import "fmt"

func ReadInput() string {
	var input string
	fmt.Println("Enter the text you want to encrypt:")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Input error, try again.")
		return ReadInput()
	}
	return input
}
