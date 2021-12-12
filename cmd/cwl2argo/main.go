package main

import (
	"cwl2argo/commands"
	"os"
	"os/exec"
)

func main() {
	err := commands.NewRootCommand().Execute()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() >= 0 {
				os.Exit(exitError.ExitCode())
			} else {
				os.Exit(137) // probably SIGTERM or SIGKILL
			}
		} else {
			println(err.Error())
			os.Exit(64)
		}
	}
}
