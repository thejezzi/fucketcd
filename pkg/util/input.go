package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserConfirmation(prompt string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s (y/n): ", prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "y" {
			return true, nil
		} else if input == "n" {
			return false, nil
		} else {
			fmt.Println("Please answer with y or n.")
		}
	}
}
