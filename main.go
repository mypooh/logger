package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

type customLogWriter struct {
	fileWriter    *os.File
	consoleWriter *os.File
}

func (w *customLogWriter) Write(p []byte) (n int, err error) {
	// Write other messages to the original writer
	if w.fileWriter != nil {
		w.fileWriter.Write(p)
		w.fileWriter.Sync()
	}
	return w.consoleWriter.Write(p)
}

func LogString(filename string, success bool, depth int, format string, value ...any) {
	var logFile *os.File
	var err error
	if filename != "" {
		logFile, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	}
	if logFile != nil {
		defer logFile.Close()
	}
	logWriter := &customLogWriter{}
	if err != nil {
		logWriter.fileWriter = nil
	} else {
		logWriter.fileWriter = logFile
	}
	logWriter.consoleWriter = os.Stdout
	prefix := "[+]"
	if !success {
		prefix = "[-]"
	}
	logger := log.New(logWriter, prefix, log.LstdFlags)
	pc, _, line, _ := runtime.Caller(depth)
	logger.Printf("Function: %s, Line: %d, Message: %s", runtime.FuncForPC(pc).Name(), line, fmt.Sprintf(format, value...))
}
