package main

import "fmt"

func main() {
	e := GenerateEnigma()

	encrypted := e.Encrypt("thefinals")
	fmt.Println(encrypted)
}
