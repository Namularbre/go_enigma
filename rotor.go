package main

import "slices"

type Rotor struct {
	input  []string
	output []string
	gap    int
	nexts  []Rotor
}

func (r *Rotor) rotate() {
	r.gap++
	if r.gap == len(r.output) {
		r.gap = 0
		r.nexts[0].rotate()
	}
}

func (r *Rotor) GetOutput(i string) string {
	if slices.Contains(r.input, i) {
		outputSize := len(r.output) - 1
		inputIndex := slices.Index(r.input, i) + r.gap
		outputIndex := inputIndex % outputSize
		output := r.output[outputIndex]

		for _, next := range r.nexts {
			output = next.GetOutput(output)
		}
		r.rotate()

		return output
	}
	return ""
}

func GenerateRotor1() *Rotor {
	changedAlphabet0 := []string{"i", "e", "o", "q", "f", "g", "k", "p", "l", "x", "y", "t", "w", "c", "r", "a", "h", "u", "j", "s", "b", "v", "m", "d", "z", "n"}
	changedAlphabet1 := []string{"e", "z", "i", "o", "q", "g", "k", "x", "l", "v", "y", "t", "f", "w", "c", "r", "a", "p", "h", "u", "j", "n", "s", "b", "m", "d"}

	return &Rotor{input: changedAlphabet0, output: changedAlphabet1, gap: 0}
}

func GenerateRotor2() *Rotor {
	changedAlphabet2 := []string{"i", "e", "q", "f", "z", "o", "p", "l", "a", "x", "g", "t", "w", "c", "y", "h", "k", "u", "j", "r", "s", "v", "m", "d", "n", "b"}
	changedAlphabet3 := []string{"c", "f", "a", "b", "g", "h", "i", "l", "d", "e", "j", "p", "s", "k", "o", "m", "y", "z", "n", "t", "q", "r", "w", "x", "u", "v"}

	return &Rotor{input: changedAlphabet2, output: changedAlphabet3, gap: 0}
}

func GenerateRotor3() *Rotor {
	changedAlphabet4 := []string{"q", "f", "g", "k", "p", "o", "x", "y", "t", "d", "l", "i", "z", "n", "w", "e", "h", "u", "j", "s", "c", "r", "a", "b", "v", "m"}
	changedAlphabet5 := []string{"l", "v", "y", "a", "z", "e", "p", "i", "o", "q", "g", "h", "u", "t", "f", "w", "c", "b", "m", "d", "r", "k", "x", "j", "n", "s"}

	return &Rotor{input: changedAlphabet4, output: changedAlphabet5, gap: 0}
}
