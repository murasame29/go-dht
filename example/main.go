package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	if err := run(); err != nil {
		log.Println("failed to run. error: ", err)
		os.Exit(1)
	}
}

func run() error {
	if err := rpio.Open(); err != nil {
		return fmt.Errorf("failed to rpio open. error: %w", err)
	}

	defer rpio.Close()

	return nil
}
