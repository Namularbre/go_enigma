package main

import (
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
	}

	if wantedOutput != output {
		t.Fatalf("Rotor error. For %v we sould have %v and not %v", input, wantedOutput, output)
	}
}

func TestRotor_GetOutputWith27CharsAnd2Rotors(t *testing.T) {
	input := []string{"i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i", "i"}
	wantedOutput := "fgchajyelropbsktdimznuqvwxa"

	r1 := GenerateRotor1()
	r2 := GenerateRotor2()
	r1.nexts = append(r1.nexts, *r2)

	var output string

	for _, char := range input {
		output += r1.GetOutput(char)
	}

	if wantedOutput != output {
		t.Fatalf("Rotor error. For %v we sould have %v and not %v", input, wantedOutput, output)
	}
}

func TestRotor_rotate1Time(t *testing.T) {
	r1 := GenerateRotor1()
	wantedGap := 1

	r1.rotate()

	if r1.gap != wantedGap {
		t.Fatalf("Rotor error. The gap after one turn sould be %v and not %v", wantedGap, r1.gap)
	}
}

func TestRotor_rotate27Time(t *testing.T) {
	r1 := GenerateRotor1()
	wantedGap := 1

	for i := 0; i < 27; i++ {
		r1.rotate()
	}

	if r1.gap != wantedGap {
		t.Fatalf("Rotor error. The gap after one turn sould be %v and not %v", wantedGap, r1.gap)
	}
}
