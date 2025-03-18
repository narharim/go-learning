package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/narharim/go-learning/concurrent-patterns/patterns"
)

func main() {
	if len(os.Args) < 2 {
		//	printUsage()
		os.Exit(1)
	}
	pattern := strings.ToLower(os.Args[1])
	switch pattern {
	case "goroutines":
		fmt.Println("=== running basic goroutines pattern ===")
		patterns.RunWithMutex()
	case "fanin":
		fmt.Println("=== running basic goroutines pattern ===")
		patterns.FanInFanOut()
	}
}
