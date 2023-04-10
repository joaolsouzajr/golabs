package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile() string {

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanLines)

	fmt.Println("Hello")

	var contentBuilder strings.Builder

	for input.Scan() {
		contentBuilder.WriteString(input.Text())
		fmt.Println(input.Text())
	}

	return contentBuilder.String()

}
