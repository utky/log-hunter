package main

import (
	"flag"
	"io/ioutil"
	"os"
)

func main() {
	configPath := flag.String(
		"config",
		"config.json",
		"file path to config json file.")
	outPath := flag.String(
		"out",
		".",
		"directory path to write output")
	logPath := flag.String(
		"log",
		"--",
		"file path to output log file")
	flag.Parse()

	configData, configErr := ioutil.ReadFile(*configPath)

	if configErr != nil {
		panic(configErr)
	}

	config, parseErr := ParseConfig(configData)

	if parseErr != nil {
		panic(parseErr)
	}

	// Build up sequece of ssh/scp commands as list of action
	sessions, err := BuildCommand(config, *outPath)

	if err != nil {
		panic(err)
	}

	//handler := NewHandler(*outPath)
	logWriter := os.Stdout
	if *logPath != "--" {
		newWriter, logWErr := os.Create(*logPath)
		if logWErr != nil {
			panic(logWErr)
		}
		logWriter = newWriter
	}
	logger := NewLogger(logWriter)

	RunSessions(sessions, logger)

}
