package log

type logger interface {
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
