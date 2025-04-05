package utils

import (
    "fmt"
    "time"
)

func FormatTimestamp(timestamp time.Time) string {
    return timestamp.Format("2006-01-02 15:04:05")
}

func PrintMessage(username, content string, timestamp time.Time) {
    fmt.Printf("[%s] %s: %s\n", FormatTimestamp(timestamp), username, content)
} 