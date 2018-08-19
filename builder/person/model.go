package person

import "github.com/lfourky/go-patterns/builder/address"

type Person struct {
	name      string
	age       int
	address   address.Address
	reallyOld bool
}

func (p Person) Name() string {
	return p.name
}

type PersonBuilder interface {
	Name(name string) PersonBuilder
	Age(age int) PersonBuilder
	Address(address address.Address) PersonBuilder
	Build() Person
}

type builder struct {
	person Person
}

// NewBuilder returns a builder which is used to further construct a Person
func NewBuilder() PersonBuilder {
	return builder{}
}

func (b builder) Name(name string) PersonBuilder {
	b.person.name = name
	return b
}

func (b builder) Age(age int) PersonBuilder {
	b.person.age = age
	if b.person.age > 90 {
		// Just to demonstrate some other logic that could be applied here (not that it's not apparent)
		b.person.reallyOld = true
	}
	return b
}

func (b builder) Address(address address.Address) PersonBuilder {
	b.person.address = address
	return b
}

// Build returns the person constructed so far
func (b builder) Build() Person {
	return b.person
}
