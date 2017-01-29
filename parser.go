package main

import (
	"encoding/json"
)

func ParseConfig(input []byte) (Config, error) {
	var config Config
	err := json.Unmarshal(input, &config)
	return config, err
}
