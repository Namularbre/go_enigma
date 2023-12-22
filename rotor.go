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
		for _, next := range r.nexts {
			next.rotate()
		}
	}
}

func (r *Rotor) GetOutput(i string) string {
	if slices.Contains(r.input, i) {

		inputIndex := slices.Index(r.input, i)
		outputIndex := (inputIndex+r.gap)%len(r.output) - 1
		output := r.output[outputIndex]

		for _, next := range r.nexts {
			output = next.GetOutput(output)
		}

		return output
	}
	return ""
}
