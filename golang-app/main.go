package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	count := 0
	for {
		count++
		fmt.Printf("[%s] STDOUT: Application heart beat count %d\n", time.Now().Format(time.RFC3339), count)

		if count%5 == 0 {
			fmt.Fprintf(os.Stderr, "[%s] STDERR: Simulating an error log entry #%d\n", time.Now().Format(time.RFC3339), count/5)
		}

		time.Sleep(2 * time.Second)
	}
}
