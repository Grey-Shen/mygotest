package loggers

type Logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
}
