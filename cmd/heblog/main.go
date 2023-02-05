package main

import (
	_ "go.uber.org/automaxprocs"
	"heblog/internal/heblog"
	"os"
)

func main() {
	command := heblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
