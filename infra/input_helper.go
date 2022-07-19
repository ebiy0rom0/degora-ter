package infra

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func WaitScan() (string, error) {
	fmt.Printf("select operation > ")
	if !scanner.Scan() {
		err := fmt.Errorf("scan error: %w", scanner.Err())
		return "", err
	}
	return scanner.Text(), nil
}
