package main

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
	for index, char := range r.input {
		if char == i {
			outputIndex := (index+r.gap)%len(r.output) - 1
			for _, n := range r.nexts {
				char = n.GetOutput(char)
			}
			r.rotate()

			return r.output[outputIndex]
		}
	}
	return ""
}
