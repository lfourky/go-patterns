package log

import "fmt"

type regular struct {
}

// New instantiates and returns a new logger implementation
func New() logger {
	return regular{}
}

func (l regular) Debugf(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("[D] %s\n", format), args...)
}

func (l regular) Errorf(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("[E] %s\n", format), args...)
}
