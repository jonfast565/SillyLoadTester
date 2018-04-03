package main

import "fmt"

func main() {
	fmt.Printf("=== Silly Load Test ===\n")
	logger := NewProgramLogger(false)
	configuration := ReadConfig("./config.json", logger)
	loadTester := InitLoadTester(configuration, logger)
	loadTester.RunTaskLoop()
}
