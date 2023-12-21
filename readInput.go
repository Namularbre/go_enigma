package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput() string {
	var i string
	fmt.Println("Enter the text you want to encrypt:")
	reader := bufio.NewReader(os.Stdin)
	i, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		fmt.Println("Input error, try again.")
		return ReadInput()
	}
	return i
}
