package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

// SetLogger sets the logger format.
func SetLoggerFormat() {
	logrus.SetFormatter(&zt_formatter.ZtFormatter{})
	logrus.SetReportCaller(true)
}
