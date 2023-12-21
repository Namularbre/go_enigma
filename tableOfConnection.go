package main

import "slices"

type TableOfConnection struct {
	input  []string
	output []string
}

func (toc *TableOfConnection) GetOutput(i string) string {
	if slices.Contains(toc.input, i) {
		for index, char := range toc.input {
			if char == i {
				return toc.output[index]
			}
		}
	}
	return ""
}
