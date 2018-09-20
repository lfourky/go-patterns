package mail

import "fmt"

type MockMailer struct {
}

// Send (in this case) mocks an actual mail sending implementation
func (m MockMailer) Send(args ...interface{}) {
	fmt.Println("Sending a mail! Content:", args)
}
