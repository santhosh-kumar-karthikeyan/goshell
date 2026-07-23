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
		commandStr := scanner.Text()
		command := strings.Split(commandStr, " ")
		switch(command[0]) {
		case "exit":
			return
		case "echo":
			fmt.Println(strings.Join(command[1:], " "))
		default:
			fmt.Printf("%s: command not found\n", command[0])
		}
	}
}
