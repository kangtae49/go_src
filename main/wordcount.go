package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	work("-")
}

func work(input string) error {
	var scanner *bufio.Scanner
	var writer *bufio.Writer

	if input == "-" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(input)
		if err != nil {
			return err
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}
	writer = bufio.NewWriter(os.Stdout)
	cnt := 0
	for scanner.Scan() {
		line := scanner.Text()
		sline := strings.TrimSpace(line)
		cols := strings.Split(sline, ",")
		out := strings.Join(cols, "-")
		//fmt.Println(out)
		fmt.Fprintln(writer, out)
		writer.Flush()
		cnt += 1
	}
	fmt.Println("line cnt: ", cnt)
	return nil
}
