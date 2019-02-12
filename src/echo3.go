// Echo1 prints ints command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var startTime, endTime time.Time
	startTime = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	endTime = time.Now()
	printElapsedTime(startTime, endTime)

	// check efficiency
	startTime = time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
	endTime = time.Now()
	printElapsedTime(startTime, endTime)

}

func printElapsedTime(t0, t1 time.Time) {
	fmt.Println("elapsed: %d", t1.Sub(t0))
}
