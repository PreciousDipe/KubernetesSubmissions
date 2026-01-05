package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	// Generate a random UUID string on startup
	randomString := uuid.New().String()
	fmt.Printf("Generated a random string: %s\n", randomString)
	fmt.Println("Outputting every 5 seconds (press Ctrl+C to stop)....")

	// Create a ticker for 5-seconds intervals
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Output the stored string with the timestap every 5 seconds
	for {
		select {
		case <-ticker.C:
			timestamp := time.Now().UTC().Format(time.RFC3339)
			fmt.Printf("%s: %s\n", timestamp, randomString)
		}
	}
}
