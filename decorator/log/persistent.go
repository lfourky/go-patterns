package log

import "github.com/lfourky/go-patterns/decorator/persistence"

type persistent struct {
	logger
	persister persistence.Persister
}

// WithPersistence decorates a logger with log persistence
func WithPersistence(l logger, p persistence.Persister) logger {
	return persistent{logger: l, persister: p}
}

func (l persistent) Debugf(format string, args ...interface{}) {
	l.persister.Persist(args...)
	l.logger.Debugf(format, args...)
}

func (l persistent) Errorf(format string, args ...interface{}) {
	l.persister.Persist(args...)
	l.logger.Errorf(format, args...)
}
