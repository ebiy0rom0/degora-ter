package main

import (
	"fmt"
	"go-terminal/infra"
	"log"
	"os"
)

func main() {
	for {
		infra.DecorateDescription(
			"input test test test test test test test test test test ",
			[]string{
				"free input",
				"test input",
				"free input",
				"test input",
			},
		)
		t, err := infra.WaitScan()
		if err != nil {
			log.Fatalf("%s", err.Error())
		} else if t == "exit" {
			os.Exit(0)
		}
		fmt.Printf("your input is: %s\n\n", t)
	}
}
