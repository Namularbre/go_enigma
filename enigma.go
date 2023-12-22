package main

import (
	"strings"
)

type Enigma struct {
	firstRotor Rotor
	toc        TableOfConnexions
}

func GenerateEnigma() *Enigma {
	var e Enigma

	toc := GenerateTableOfConnexions()

	r1 := GenerateRotor1()
	r2 := GenerateRotor2()
	r3 := GenerateRotor3()

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
