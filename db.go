package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData() (MoneyData, error) {
	filename := os.Getenv("DB_FILE")
	if filename == "" {
		return MoneyData{}, fmt.Errorf("DB_FILE env var not set")
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return MoneyData{}, err
	}

	split := strings.Split(string(file), "\n")

	money := split[0]

	lastUpdate, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return MoneyData{}, err
	}

	nextUpdate, err := strconv.ParseInt(split[2], 10, 64)
	if err != nil {
		return MoneyData{}, err
	}

	return MoneyData{lastUpdate, money, nextUpdate}, nil
}
