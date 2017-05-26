package postgres

import "usefulinterfaces/exercises/impl"

type CustomerService struct{}

func (c *CustomerService) Get(id int) (*impl.Customer, error) {
	panic("not implemented")
}
