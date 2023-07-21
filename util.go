package main

import (
	"fmt"
	"time"
)

func generateDynamicFilename(baseFilename, urlNumber, fileExt string) string {
	// Use the current timestamp as a unique identifier
	timestamp := time.Now().UnixNano()

	// Concatenate the base filename, timestamp, and file extension
	filename := fmt.Sprintf("data/%s_%s_%d.%s", baseFilename, urlNumber, timestamp, fileExt)
	return filename
}