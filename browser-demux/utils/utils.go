package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ConfirmAction(prompt string) (bool, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		return false, err
	}
	return char == 'y', nil
}
