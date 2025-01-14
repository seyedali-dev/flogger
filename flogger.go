package flogger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"path"
)

// log is a global logger instance that will be used throughout the application.
var log *logrus.Logger

// customFormatter is a custom log formatter that extends the prefixed.TextFormatter.
// It adds additional fields like function name and file location to the log output.
type customFormatter struct {
	*prefixed.TextFormatter
}

// Format is a method that overrides the default Format method of logrus.Entry.
// It adds custom fields (function name and file location) to the log entry if the caller information is available.
func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Check if the log entry has caller information (file and line number).
	if entry.HasCaller() {
		// Extract the function name from the caller.
		funcVal := entry.Caller.Function
		// Extract the file name and line number from the caller and format it as "file:line".
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		// Initialize the log entry's data fields if they are nil.
		if entry.Data == nil {
			entry.Data = make(logrus.Fields)
		}

		// Add the function name and file location to the log entry's data.
		entry.Data["func"] = funcVal
		entry.Data["file"] = fileVal
	}

	// Use the parent TextFormatter to format the log entry.
	return f.TextFormatter.Format(entry)
}

// init is a special function that initializes the logger when the package is imported.
func init() {
	// Create a new instance of the logrus logger.
	log = logrus.New()

	// Initialize the custom formatter with desired settings.
	formatter := &customFormatter{
		TextFormatter: &prefixed.TextFormatter{
			ForceColors:     true,                  // Force colored output.
			ForceFormatting: true,                  // Force formatting even if the output is not a terminal.
			FullTimestamp:   true,                  // Include the full timestamp in the log output.
			TimestampFormat: "2006-01-02 15:04:05", // Set the timestamp format.
		},
	}

	// Set the custom formatter as the logger's formatter.
	log.SetFormatter(formatter)

	// Uncomment the following line to enable caller information (file and line number) in logs.
	// log.SetReportCaller(true)

	// Set the default log level to Info. Adjust this as needed for your application.
	log.SetLevel(logrus.InfoLevel)
}

// log level functions

// Info logs a message at the Info level with formatting.
// It accepts a format string and variadic arguments, similar to fmt.Printf.
func Info(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warn logs a message at the Warn level with formatting.
// It accepts a format string and variadic arguments, similar to fmt.Printf.
func Warn(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Error logs a message at the Error level with formatting.
// It accepts a format string and variadic arguments, similar to fmt.Printf.
func Error(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
