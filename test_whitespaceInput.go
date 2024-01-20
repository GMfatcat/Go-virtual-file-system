package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func extractWhiteSpaceInput(input string) ([]string, error) {
	emptyStringList := []string{}

	var parts []string
	var currentPart string
	var insideQuotes bool = false

	for _, char := range input {
		if char == ' ' && !insideQuotes {
			if currentPart != "" {
				parts = append(parts, currentPart)
				currentPart = ""
			}
		} else if char == '"' {
			insideQuotes = !insideQuotes
		} else {
			currentPart += string(char)
		}
	}

	if currentPart != "" {
		parts = append(parts, currentPart)
	}

	if insideQuotes {
		return emptyStringList, fmt.Errorf("unclosed double quote")
	}

	if len(parts) == 0 {
		return emptyStringList, fmt.Errorf("no keywords provided")
	}

	if len(parts) == 1 {
		return emptyStringList, fmt.Errorf("no arguments provided")
	}

	return parts, nil
}

func RegexCheck(input string) error {
	inputRegex, err := regexp.Compile(`^[A-Za-z0-9 ]{3,10}$`)
	if err != nil {
		return fmt.Errorf("Regex Compile Error: %v\n", err)
	}

	if !inputRegex.MatchString(input) {
		return fmt.Errorf("The %s contain invalid chars.\n", input)
	}
	return nil
}

func main() {
	for {

		fmt.Print("Enter command: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if input == "exit" {
			return
		}

		parts, err := extractInput(input)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Parts:", parts, "Num:", len(parts))

		for _, part := range parts {
			if err := RegexCheck(part); err != nil {
				fmt.Println("Error:", err)
				return
			}
		}
	}
}
