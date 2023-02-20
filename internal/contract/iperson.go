package contract

type IPerson interface {
	Say() string
}

type Person struct {
}

func (p Person) Say() string {
	return "my name is chentenglong"
}
