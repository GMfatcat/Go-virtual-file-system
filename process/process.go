/* Process User Input Commands */
package process

import (
	"fmt"
	"strings"
)

/*Example: register username*/
func ProcessInput(command string) error {

	return nil
}

// 1 keyword + at least 1 args
func extractInput(input string) ([]string, error) {

	emptyStringList := []string{}

	parts := strings.Fields(input)

	if len(parts) == 0 {
		return emptyStringList, fmt.Errorf("No keyword found in the input")
	}

	if len(parts) == 1 {
		return emptyStringList, fmt.Errorf("No keyword argument found in the input")
	}

	// keyword := parts[0]

	return parts, nil
}
