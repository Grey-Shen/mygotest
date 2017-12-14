package liblogger

import (
	"github.com/sirupsen/logrus"
	"gitlab.qiniu.io/pingan/libbase/loggers"
)

type Logger interface {
	loggers.Logger
	WithError(error) *logrus.Entry
	WithField(string, interface{}) *logrus.Entry
	WithFields(logrus.Fields) *logrus.Entry
}
