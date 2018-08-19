package address

type Address struct {
	street string
	number int
}

type AddressBuilder interface {
	Street(street string) AddressBuilder
	Number(number int) AddressBuilder
	Build() Address
}

type builder struct {
	address Address
}

// NewBuilder returns a builder which is used to further construct an Address
func NewBuilder() AddressBuilder {
	return builder{}
}

func (b builder) Builder() AddressBuilder {
	return builder{}
}

func (b builder) Street(street string) AddressBuilder {
	b.address.street = street
	return b
}

func (b builder) Number(number int) AddressBuilder {
	b.address.number = number
	return b
}

// Build returns the address constructed so far
func (b builder) Build() Address {
	return b.address
}
