package mail

type Sender interface {
	Send(args ...interface{})
}
