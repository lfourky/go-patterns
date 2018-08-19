package main

import (
	"fmt"

	"github.com/lfourky/go-patterns/builder/address"
	"github.com/lfourky/go-patterns/builder/email"
	"github.com/lfourky/go-patterns/builder/person"
)

func main() {
	p1 := person.NewBuilder().
		Name("Lazar Ivkov").
		Age(25).
		Address(address.NewBuilder().
			Street("The Avenue des Champs-Élysées").
			Number(42).
			Build()).
		Build()

	fmt.Printf("%#v\n", p1)

	mail, errors := email.NewBuilder().
		Recipient("lfourky@gmail.com").
		Subject("disturbance in the force").
		Body("").
		Sender(p1).
		Build()

	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
		//optional return or error handling
	}

	fmt.Printf("%#v\n", mail)

	if err := mail.Send(); err != nil {
		fmt.Println(err)
	}
}
