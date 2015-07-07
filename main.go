package main

import (
	"bufio"
	"fmt"
	"github.com/shiena/ansicolor"
	"os"
)

func main() {
	w := ansicolor.NewAnsiColorWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintln(w, scanner.Text())
	}
}
