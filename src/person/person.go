package person

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(name string) {
	p.firstName = name
}
