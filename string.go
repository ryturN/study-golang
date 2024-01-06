package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	counter := 10
	for counter > 0 {
		fmt.Println(counter)
		counter--
	}
	startTime := time.Now()
	for counter := 100; counter >= 0; counter-- {
		fmt.Println("perulangan ke 	\t", counter)
	}
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Println("waktu yang diperlukan", elapsed)
	fmt.Println("done")

	start := time.Now()
	var output strings.Builder

	for counter := 100; counter >= 0; counter-- {
		if counter%2 == 1 {
			continue
		}
		fmt.Fprintln(&output, "perulangan ke \t", counter)
	}
	fmt.Println(output.String())
	end := time.Now()
	elapsed = end.Sub(start)
	fmt.Println("waktu yang diperlukan", elapsed)
}
