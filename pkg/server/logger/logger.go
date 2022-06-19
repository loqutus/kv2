package logger

import (
	"github.com/sirupsen/logrus"
)

// SetLogger sets the logger format.
func SetLoggerFormat() {
	// logrus.SetFormatter(&zt_formatter.ZtFormatter{})
	logrus.SetReportCaller(true)
}
