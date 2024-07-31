package main

import (
	"fmt"
	"testing"
)

func Test_Clean_Input(t *testing.T) {

	testInputs := struct {
		inputs  []string
		outputs []string
	}{
		inputs:  []string{"", "wHaT", "I nnoooooooooooo watch outttt"},
		outputs: []string{"", "what", "i"},
	}

	clean_in := []string{}
	for _, input := range testInputs.inputs {

		// How to check if empty string[]
		if len(input) > 0 {
			res := cleanInput(input)[0]
			clean_in = append(clean_in, res)
			continue
		} else {
			clean_in = append(clean_in, "")
			continue
		}

	}

	fmt.Println()
	length_clean := len(clean_in)
	fmt.Printf("Length of clean in is %d\n", length_clean)
	fmt.Println()

	for n, in := range clean_in {
		if in != testInputs.outputs[n] {
			t.Errorf("Error: Cleaned Input: %s does not match Output: %s \n", in, testInputs.outputs[n])
		}
	}

}
