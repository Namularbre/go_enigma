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
