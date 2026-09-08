package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	start := time.Now()

	var med string
	for i := 0; i < len(os.Args); i++ {
		med = med + " " + os.Args[i]
	}

	sec := time.Since(start)
	fmt.Println("Медленно:", sec)

	start2 := time.Now()

	fmt.Println(strings.Join(os.Args, " "))

	sec1 := time.Since(start2)
	fmt.Println("быстро:", sec1)
}