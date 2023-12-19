package main

type TableOfConnection struct {
	input  []string
	output []string
}

func (toc *TableOfConnection) GetOutput(i string) string {
	for index, char := range toc.input {
		if char == i {
			return toc.output[index]
		}
	}
	return ""
}
