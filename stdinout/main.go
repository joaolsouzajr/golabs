package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile()
}

func readCommand() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello")

	for true {
		fmt.Print("Type command: \n")
		input.Scan()
		name := input.Text()
		if name == "Exit" {
			fmt.Printf("%s, God Bay\n", name)
			break
		}
	}
}
