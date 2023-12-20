package main

import "fmt"

func main() {
	e := GenerateEnigma()

	in := ReadInput()

	encrypted := e.Encrypt(in)
	fmt.Println(encrypted)
}
