package main

import (
	"fmt"
	"go.uber.org/zap"
	"runtime"
)

type ProgramLogger struct {
	IsProduction bool
	Logger       *zap.Logger
}

func NewProgramLogger(isProduction bool) *ProgramLogger {
	logger, _ := zap.NewProduction()
	return &ProgramLogger{
		isProduction,
		logger,
	}
}

func (logger *ProgramLogger) LogError(error string) {
	if logger.IsProduction {
		logger.Logger.Error(error)
	} else {
		fmt.Printf("Error: " + error + getNewLineCharacter())
	}
}

func (logger *ProgramLogger) LogSuccess(success string) {
	if logger.IsProduction {
		logger.Logger.Info(success)
	} else {
		fmt.Printf(success + getNewLineCharacter())
	}
}

func getNewLineCharacter() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}