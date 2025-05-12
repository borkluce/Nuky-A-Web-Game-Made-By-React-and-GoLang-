package env

import (
	// Standart
	"os"

	// Package specific
	"github.com/kahleryasla/pkg/go/log/util"
)

func LogWorkingEnv() {
	// Log where is the working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		util.LogError("failed to get working directory", "LogWorkingEnv()", "")
		return
	}
	util.LogInfo("working directory: "+workingDirectory, "LogWorkingEnv()", "")

	nowDir, err := os.ReadDir(".")
	if err != nil {
		util.LogError("failed to read current directory", "LogWorkingEnv()", "")
		return
	}
	name := nowDir[0].Name()
	util.LogInfo("current directory: "+name, "LogWorkingEnv()", "")
}
