package services

import "github.com/xm-chentl/goframedemo/internal/contract"

type Person struct {
}

func (p Person) Say() string {
	return "my name is chentenglong"
}

func NewPersonService() contract.IPersonService {
	return &Person{}
}
