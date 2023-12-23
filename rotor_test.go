package main

import (
	"fmt"
	"testing"
)

func TestRotor_GetOutputWith1Char(t *testing.T) {
	input := "a"
	wantedOutput := "r"

	r1 := GenerateRotor1()
	output := r1.GetOutput(input)

	if wantedOutput != output {
		t.Fatalf("Rotor error. For %v we sould have %v and not %v", input, wantedOutput, output)
	}
}

func TestRotor_GetOutputUnknownChar(t *testing.T) {
	input := "?"
	wantedOutput := ""

	r1 := GenerateRotor1()
	output := r1.GetOutput(input)

	if wantedOutput != output {
		t.Fatalf("Rotor error. For %v we sould have %v and not %v", input, wantedOutput, output)
	}
}

func TestRotor_GetOutputWith5Chars(t *testing.T) {
	input := []string{"t", "e", "s", "t", "s"}
	wantedOutput := "tincb"

	r1 := GenerateRotor1()

	var output string

	for _, char := range input {
		output += r1.GetOutput(char)
	}

	if wantedOutput != output {
		t.Fatalf("Rotor error. For %v we sould have %v and not %v", input, wantedOutput, output)
	}
}

func TestRotor_GetOutputWith1CharAnd3Rotors(t *testing.T) {
	input := "a"
	wantedOutput := "o"

	r1 := GenerateRotor1()
	r2 := GenerateRotor2()
	r3 := GenerateRotor3()
	r1.nexts = []Rotor{*r2, *r3}
	output := r1.GetOutput(input)

	if wantedOutput != output {
		t.Fatalf("Rotor error. For %v we sould have %v and not %v", input, wantedOutput, output)
	}
}

func TestRotor_GetOutputWith27(t *testing.T) {
	input := []string{"i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i"}
	wantedOutput := "ezioqgkxlvytfwcraphujnsbmde"
	r1 := GenerateRotor1()
	var output string

	for _, char := range input {
		output += r1.GetOutput(char)
		fmt.Println(len(output))
	}

	if wantedOutput != output {
		t.Fatalf("Rotor error. For %v we sould have %v and not %v", input, wantedOutput, output)
	}
}
