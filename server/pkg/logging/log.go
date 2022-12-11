package logging

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/ttys3/rotatefilehook"
)

var (
	DefaultCallerDepth = 2

	logger *logrus.Logger = logrus.New()
)

// Setup initialize the log instance
func Setup() {
	filePath := getLogFilePath()
	fileName := getLogFileName()
	logger = logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		PadLevelText:     true,
		DisableSorting:   true,
	})

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   fmt.Sprintf("%s/%s", filePath, fileName),
		MaxSize:    10, // the maximum size in megabytes
		MaxBackups: 5,  // the maximum number of old log files to retain
		MaxAge:     7,  // the maximum number of days to retain old log files
		LocalTime:  true,
		Level:      logrus.InfoLevel,
		Formatter:  &logrus.JSONFormatter{},
	})
	if err != nil {
		logger.Fatal(err)
	}

	logger.AddHook(rotateFileHook)
}

// Info output logs at info level
func Debugln(v ...interface{}) {
	addFields().Debugln(v...)
}

// Info output logs at info level
func Infoln(args ...interface{}) {
	addFields().Infoln(args...)
}

// Infof output logs at info level
func Infof(s string, v ...interface{}) {
	addFields().Infof(s, v...)
}

// Error output logs at error level
func Errorln(v ...interface{}) {
	addFields().Errorln(v...)
}

// Error output logs at error level
func Errorf(format string, args ...interface{}) {
	addFields().Errorf(format, args...)
}

// Fatal output logs at fatal level
func Fatalln(v ...interface{}) {
	addFields().Fatalln(v...)
}

// setPrefix set the prefix of the log output
func addFields() *logrus.Entry {
	fields := map[string]interface{}{}

	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		base := filepath.Base(filepath.Dir(file))
		file = filepath.Base(file)

		fields = logrus.Fields{
			"caller": fmt.Sprintf("%s/%s:%d", base, file, line),
		}
	}

	return logger.WithFields(fields)
}
