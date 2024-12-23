package main

import (
	"log"
	"os"
)

func main() {
	if err := validateInput(os.Args); err != nil {
		log.Fatal(err)
	}

	processName := os.Args[1]

	err := findAndKillProcess(processName)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Println("Done.")
	}
}
