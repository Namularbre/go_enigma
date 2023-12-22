package main

import "slices"

type TableOfConnexions struct {
	input  []string
	output []string
}

func (toc *TableOfConnexions) GetOutput(i string) string {
	if slices.Contains(toc.input, i) {
		inputIndex := slices.Index(toc.input, i)
		return toc.output[inputIndex]
	}
	return ""
}

func GenerateTableOfConnexions() *TableOfConnexions {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	changedAlphabet := []string{"a", "c", "i", "e", "f", "g", "h", "j", "k", "l", "x", "y", "o", "w", "p", "q", "r", "d", "s", "b", "t", "u", "v", "m", "z", "n"}
	return &TableOfConnexions{input: alphabet, output: changedAlphabet}
}
