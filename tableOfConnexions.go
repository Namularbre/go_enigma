package main

import "slices"

type TableOfConnexions struct {
	input  []string
	output []string
}

func (toc *TableOfConnexions) GetOutput(i string) string {
	if slices.Contains(toc.input, i) {
		for index, char := range toc.input {
			if char == i {
				return toc.output[index]
			}
		}
	}
	return ""
}

func (toc *TableOfConnexions) GetInput(o string) string {
	if slices.Contains(toc.output, o) {
		for index, char := range toc.output {
			if char == o {
				return toc.input[index]
			}
		}
	}
	return ""
}
