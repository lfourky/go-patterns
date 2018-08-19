package email

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lfourky/go-patterns/builder/person"
)

type email struct {
	recipient string
	subject   string
	body      string
	sender    person.Person
}

func (e email) Send() error {
	return errors.New("currently depends heavily on your imagination")
}

type EmailBuilder interface {
	Recipient(address string) EmailBuilder
	Subject(subject string) EmailBuilder
	Body(body string) EmailBuilder
	Sender(p person.Person) EmailBuilder
	Build() (email, []error)
}

type builder struct {
	email email
}

// NewBuilder returns a builder which is used to further construct an Email
func NewBuilder() EmailBuilder {
	return builder{}
}

func (b builder) Recipient(address string) EmailBuilder {
	b.email.recipient = address
	return b
}

func (b builder) Subject(subject string) EmailBuilder {
	b.email.subject = subject
	return b
}

func (b builder) Body(body string) EmailBuilder {
	b.email.body = body
	return b
}

func (b builder) Sender(p person.Person) EmailBuilder {
	b.email.sender = p
	return b
}

// Build returns the Email constructed so far, and errors if some fields didn't pass validation
func (b builder) Build() (email, []error) {
	var validationErrors = make([]error, 0)

	if strings.Contains(b.email.recipient, "@") {
		//Quite a rigorous email validation going on over here
		validationErrors = append(validationErrors, errors.New("recipient address should have '@' symbol"))
	}

	if b.email.recipient == "lfourky@gmail.com" {
		validationErrors = append(validationErrors, errors.New("please do not disturb the almighty creator"))
	}

	if b.email.sender.Name() == "Lazar Ivkov" {
		validationErrors = append(validationErrors, errors.New("now that's quite lonely tbh"))
	}

	//etc.

	//Add that 'kind regards' sentence at the end, we're being polite over here
	b.email.body += fmt.Sprintf("\n Kind regards,\n%s", b.email.sender.Name())

	return b.email, validationErrors
}
