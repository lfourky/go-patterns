package main

import (
	"fmt"

	"github.com/lfourky/go-patterns/builder/address"
	"github.com/lfourky/go-patterns/builder/person"
)

func main() {
	p1 := person.NewBuilder().
		Name("Lazar Ivkov").
		Age(25).
		Address(address.NewBuilder().
			Street("The Avenue des Champs-Élysées").
			Number(42).
			Build())
	fmt.Printf("%#v\n", p1)
}
