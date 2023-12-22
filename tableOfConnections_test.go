package main

import (
	"testing"
)

func TestTableOfConnexions_GetOutputExistingChar(t *testing.T) {
	input := "t"
	wantedOutput := "b"

	toc := GenerateTableOfConnexions()
	output := toc.GetOutput(input)

	if wantedOutput != output {
		t.Fatalf("Table of connections error. For %v we sould have %v and not %v", input, wantedOutput, output)
	}
}

func TestTableOfConnexions_GetOutputNotExistingChar(t *testing.T) {
	input := "😎"
	wantedOutput := ""

	toc := GenerateTableOfConnexions()
	output := toc.GetOutput(input)

	if wantedOutput != output {
		t.Fatalf("Table of connections error. For %v we sould have %v and not %v", input, wantedOutput, output)
	}
}
