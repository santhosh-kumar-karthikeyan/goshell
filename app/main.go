package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	for {
		fmt.Print("$ ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := scanner.Text()
		if strings.Compare("exit", command) == 0 {
			return
		}
		fmt.Printf("%s: command not found\n", command)
	}
}
