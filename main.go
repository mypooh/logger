package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

type customLogWriter struct {
	// this is file to write log
	fileWriter *os.File
	// this is std
	consoleWriter *os.File
}

func (w *customLogWriter) Write(p []byte) (n int, err error) {
	// Write other messages to the original writer
	if w.fileWriter != nil {
		w.fileWriter.Seek(0, 2)
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
	logger.Printf("Function: %s, Line: %d, Message: %s\n", runtime.FuncForPC(pc).Name(), line, fmt.Sprintf(format, value...))
}

func FullLogString(filename string, success bool, depth int, format string, value ...any) {
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
	logger.Printf("Function: %s, Line: %d, Message: %s\n", runtime.FuncForPC(pc).Name(), line, fmt.Sprintf(format, value...))
}
