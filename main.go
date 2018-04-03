package main

import "fmt"

func main() {
	fmt.Printf("=== Silly Load Test ===" + getNewLineCharacter())
	logger := NewProgramLogger(false)
	configuration := ReadConfig("./config.json", logger)
	loadTester := InitLoadTester(configuration, logger)
	loadTester.RunTaskLoop()
	duration := loadTester.GetAverageDuration()
	fmt.Printf("Average time per all requests: %s", duration)
}
