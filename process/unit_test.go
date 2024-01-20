/* Unit test for Process package*/
package process

import (
	"fmt"
	"testing"
)

func TestExtractInput(t *testing.T) {

	t.Skip()

	var correct = 0
	var incorrect = 0

	// Valid Test Case: keyword Arg1 Arg2...
	// Last 2 strings are invalid
	inputs := []string{
		"register user1",
		"create-file filename.txt",
		"some-other-command argument1 argument2",
		"some-other-command --argument1 argument2",
		"register",
		"create-file",
	}

	for _, input := range inputs {
		if _, err := extractInput(input); err != nil {
			incorrect++
		} else {
			correct++
		}
	}

	if incorrect == 2 && correct == 4 {
		t.Log("Expected 2 incorrect and 4 correct, got 2 incorrect and 4 correct")
	} else {
		t.Errorf("Expected 2 incorrect and 4 correct, got %d incorrect and %d correct",
			incorrect, correct)
	}

}

func TestExtractWhiteSpaceInput(t *testing.T) {

	t.Skip()

	var testString = []string{
		"register fatcat",
		`register "fat cat"`,
		"some one select",
		`"register" the "fat cat"`}

	var expected = []int{2, 2, 3, 3}

	// test string extraction
	for i, str := range testString {
		parts, _ := extractWhiteSpaceInput(str)
		if len(parts) != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], len(parts))
			fmt.Println("Failed Parts:", parts)
			continue
		}

		for j, word := range parts {
			fmt.Printf("%d %s\n", j, word)
		}
		fmt.Printf("========\n")

	}
}
