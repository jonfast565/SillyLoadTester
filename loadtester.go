package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LoadTester struct {
	config        *Configuration
	programLogger *ProgramLogger
}

func InitLoadTester(config *Configuration, programLogger *ProgramLogger) *LoadTester {
	return &LoadTester{
		config,
		programLogger,
	}
}

func (loadTester *LoadTester) RunTaskLoop() {
	var wg sync.WaitGroup
	wg.Add(loadTester.config.NumberOfTasks)
	for i := 0; i < loadTester.config.NumberOfTasks; i++ {
		go func(currentTask int) {
			loadTester.RunCallLoop(currentTask)
			wg.Done()
		}(i)
	}
	wg.Wait()
	loadTester.programLogger.LogSuccess("Run completed!")
}

func (loadTester *LoadTester) RunCallLoop(task int) {
	for i := 0; i < loadTester.config.CallsPerTask; i++ {
		loadTester.RunOnce(task, i)
	}
}

func (loadTester *LoadTester) RunOnce(task int, call int) {
	var resp *http.Response
	var err error

	reqString := fmt.Sprintf("Task %d, Call %d on %s", task, call, loadTester.config.RequestUrl)
	loadTester.programLogger.LogSuccess(reqString)
	start := time.Now()

	if loadTester.config.RequestVerb == "GET" {
		resp, err = http.Get(loadTester.config.RequestUrl)
	} else if loadTester.config.RequestVerb == "POST" {
		var jsonString = []byte(loadTester.config.PostBody)
		resp, err = http.Post(
			loadTester.config.RequestUrl,
			loadTester.config.PostContentType,
			bytes.NewBuffer(jsonString))
	} else {
		panic(fmt.Sprintf("Unimplemented verb %s in configuration", loadTester.config.RequestVerb))
	}

	if err != nil {
		loadTester.programLogger.LogError(err.Error())
		panic(err)
	}

	if resp.StatusCode >= 400 {
		errText := fmt.Sprintf("Task %d, Call %d Response code was %d, error status %s",
			task, call, resp.StatusCode, resp.Status)
		loadTester.programLogger.LogError(errText)
	} else {
		successText := fmt.Sprintf("Task %d, Call %d Response code was %d, status %s",
			task, call, resp.StatusCode, resp.Status)
		loadTester.programLogger.LogSuccess(successText)
	}

	elapsed := time.Since(start)
	elapsedMessage := fmt.Sprintf("Task %d, Call %d took %s", task, call, elapsed)
	loadTester.programLogger.LogSuccess(elapsedMessage)
}
