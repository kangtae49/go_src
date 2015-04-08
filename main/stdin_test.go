package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Split(line, ",")
		out := strings.Join(cols, "-")
		fmt.Println(out) // or do something else with line
	}
}
