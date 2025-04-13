package main

import (
	"fmt"
	"time"
)

func addLog(format string, args ...interface{}) {
	logsMutex.Lock()
	defer logsMutex.Unlock()
	logMessage := fmt.Sprintf(format, args...)
	logs = append(logs, logMessage)
}

func addLogWithTimestamp(format string, args ...interface{}) {
	logsMutex.Lock()
	defer logsMutex.Unlock()
	timestamp := time.Now().Format(time.RFC1123)
	logMessage := fmt.Sprintf("[%s] "+format, append([]interface{}{timestamp}, args...)...)
	logs = append(logs, logMessage)
}
