package logger

import (
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type JSONFileLogger struct {
	log *logrus.Logger
}

func NewJSONFileLogger(filePath string) (*JSONFileLogger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})

	return &JSONFileLogger{log: logger}, nil
}

func (j *JSONFileLogger) Error(message string, err error, fields map[string]interface{}) {
	errorFields := make(map[string]interface{})
	errorFields["error"] = err.Error()

	// Add any additional fields to the error log
	for key, value := range fields {
		errorFields[key] = value
	}

	j.log.WithFields(errorFields).Error(message)
}

func (j *JSONFileLogger) Info(message string, fields map[string]interface{}) {
	if fields != nil {
		j.log.WithFields(fields).Info(message)
	} else {
		j.log.Info(message)
	}
}

func ErrorLogger(l *JSONFileLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Continue to the next middleware or route handler
		c.Next()
		path := c.Request.URL.Path
		pathSegments := strings.Split(path, "/")
		service := pathSegments[1]

		// Check if any errors occurred during the request handling
		err := c.Errors.Last()

		file, exists := c.Get("file")
		if !exists {
			file = ""
		}

		line, exists := c.Get("line")
		if !exists {
			line = ""
		}

		if err != nil {
			// Get the custom message from the context (if provided)

			l.log.WithField("status", c.Writer.Status()).
				WithField("method", c.Request.Method).
				WithField("path", c.Request.URL.Path).
				WithField("service", service).
				WithField("file", file).
				WithField("line", line).
				Error(err.Err)

		}

		if c.Writer.Status() == 200 {
			l.log.WithField("status", c.Writer.Status()).
				WithField("method", c.Request.Method).
				WithField("path", c.Request.URL.Path).
				WithField("service", service).
				WithField("file", file).
				WithField("line", line).
				Info("Request completed.")
		}

	}
}

func getLineNumber() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

// Helper function to get the file name where the function is called
func getFileName() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}
