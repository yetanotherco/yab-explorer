package configs

import (
	"os"
	"strings"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func InitLog() {
	log.SetLevel(getLoggerLevel(os.Getenv("LOGGING_LEVEL")))
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
		CallerFirst:     true,
	})
	log.SetOutput(os.Stdout)
}

func getLoggerLevel(level string) log.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return log.DebugLevel
	case "TRACE":
		return log.TraceLevel
	default:
		return log.InfoLevel
	}
}
