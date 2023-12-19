package main

import "fmt"

type Enigma struct {
	firstRotor Rotor
	toc        TableOfConnection
}

func GenerateEnigma() *Enigma {
	e := Enigma{}

	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	changedAlphabet := []string{"a", "c", "i", "e", "f", "g", "h", "j", "k", "l", "x", "y", "o", "w", "p", "q", "r", "d", "s", "b", "t", "u", "v", "m", "z", "n"}
	toc := TableOfConnection{input: alphabet, output: changedAlphabet}

	changedAlphabet0 := []string{"i", "e", "o", "q", "f", "g", "k", "p", "l", "x", "y", "t", "w", "c", "r", "a", "h", "u", "j", "s", "b", "v", "m", "d", "z", "n"}
	changedAlphabet1 := []string{"e", "z", "i", "o", "q", "g", "k", "x", "l", "v", "y", "t", "f", "w", "c", "r", "a", "p", "h", "u", "j", "n", "s", "b", "m", "d"}

	changedAlphabet2 := []string{"i", "e", "q", "f", "z", "o", "p", "l", "a", "x", "g", "t", "w", "c", "y", "h", "k", "u", "j", "r", "s", "v", "m", "d", "n", "b"}
	changedAlphabet3 := []string{"c", "f", "a", "b", "g", "h", "i", "l", "d", "e", "j", "p", "s", "k", "o", "m", "y", "z", "n", "t", "q", "r", "w", "x", "u", "v"}

	changedAlphabet4 := []string{"q", "f", "g", "k", "p", "o", "x", "y", "t", "d", "l", "i", "z", "n", "w", "e", "h", "u", "j", "s", "c", "r", "a", "b", "v", "m"}
	changedAlphabet5 := []string{"l", "v", "y", "a", "z", "e", "p", "i", "o", "q", "g", "h", "u", "t", "f", "w", "c", "b", "m", "d", "r", "k", "x", "j", "n", "s"}

	r1 := Rotor{input: changedAlphabet0, output: changedAlphabet1, gap: 0}
	r2 := Rotor{input: changedAlphabet2, output: changedAlphabet3, gap: 0}
	r3 := Rotor{input: changedAlphabet4, output: changedAlphabet5, gap: 0}

	r1.nexts = append(r1.nexts, r2)
	r1.nexts = append(r1.nexts, r3)

	e.toc = toc
	e.firstRotor = r1

	return &e
}

func (e *Enigma) Encrypt(plain string) string {
	var encrypted string

	for _, cint := range plain {
		c := string(cint)
		c = e.toc.GetOutput(c)

		if c == "" {
			fmt.Println("Error in the table of connection")
		}
		c = e.firstRotor.GetOutput(c)

		if c == "" {
			fmt.Println("Error in the rotor")
		}
		encrypted += c
	}

	return encrypted
}
