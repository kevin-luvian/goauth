package logging

import (
	"fmt"
	"time"

	"github.com/kevin-luvian/goauth/server/pkg/setting"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	s := setting.App

	return fmt.Sprintf("%s%s", s.BinaryRootPath, s.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	s := setting.App

	return fmt.Sprintf("%s%s.%s",
		s.LogSaveName,
		time.Now().Format(s.TimeFormat),
		s.LogFileExt,
	)
}
