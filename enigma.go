package main

import (
	"strings"
)

type Enigma struct {
	firstRotor Rotor
	toc        TableOfConnexions
}

func generateTableOfConnexions() *TableOfConnexions {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	changedAlphabet := []string{"a", "c", "i", "e", "f", "g", "h", "j", "k", "l", "x", "y", "o", "w", "p", "q", "r", "d", "s", "b", "t", "u", "v", "m", "z", "n"}
	return &TableOfConnexions{input: alphabet, output: changedAlphabet}
}

func generateRotor1() *Rotor {
	changedAlphabet0 := []string{"i", "e", "o", "q", "f", "g", "k", "p", "l", "x", "y", "t", "w", "c", "r", "a", "h", "u", "j", "s", "b", "v", "m", "d", "z", "n"}
	changedAlphabet1 := []string{"e", "z", "i", "o", "q", "g", "k", "x", "l", "v", "y", "t", "f", "w", "c", "r", "a", "p", "h", "u", "j", "n", "s", "b", "m", "d"}

	return &Rotor{input: changedAlphabet0, output: changedAlphabet1, gap: 0}
}

func generateRotor2() *Rotor {
	changedAlphabet2 := []string{"i", "e", "q", "f", "z", "o", "p", "l", "a", "x", "g", "t", "w", "c", "y", "h", "k", "u", "j", "r", "s", "v", "m", "d", "n", "b"}
	changedAlphabet3 := []string{"c", "f", "a", "b", "g", "h", "i", "l", "d", "e", "j", "p", "s", "k", "o", "m", "y", "z", "n", "t", "q", "r", "w", "x", "u", "v"}

	return &Rotor{input: changedAlphabet2, output: changedAlphabet3, gap: 0}
}

func generateRotor3() *Rotor {
	changedAlphabet4 := []string{"q", "f", "g", "k", "p", "o", "x", "y", "t", "d", "l", "i", "z", "n", "w", "e", "h", "u", "j", "s", "c", "r", "a", "b", "v", "m"}
	changedAlphabet5 := []string{"l", "v", "y", "a", "z", "e", "p", "i", "o", "q", "g", "h", "u", "t", "f", "w", "c", "b", "m", "d", "r", "k", "x", "j", "n", "s"}

	return &Rotor{input: changedAlphabet4, output: changedAlphabet5, gap: 0}
}

func GenerateEnigma() *Enigma {
	var e Enigma

	toc := generateTableOfConnexions()

	r1 := generateRotor1()
	r2 := generateRotor2()
	r3 := generateRotor3()

	r1.nexts = append(r1.nexts, *r2)
	r1.nexts = append(r1.nexts, *r3)

	e.toc = *toc
	e.firstRotor = *r1

	return &e
}

func (e *Enigma) Encrypt(plain string) string {
	plain = strings.ToLower(plain)
	var encrypted string

	for _, cint := range plain {
		c := string(cint)
		c = e.toc.GetOutput(c)
		c = e.firstRotor.GetOutput(c)
		c = e.toc.GetOutput(c)

		encrypted += c
	}

	return encrypted
}
