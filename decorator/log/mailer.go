package log

import (
	"github.com/lfourky/go-patterns/decorator/mail"
)

type mailer struct {
	logger
	mailService mail.Sender
}

// WithMailAlerts decorates a logger with mail sending ability
func WithMailAlerts(l logger, s mail.Sender) logger {
	return mailer{logger: l, mailService: s}
}

// Override only the method you need from the initial logger (or all the methods, if that's what you're after)
func (l mailer) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
	l.mailService.Send(args...)
}
