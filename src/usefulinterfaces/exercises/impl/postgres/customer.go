package postgres

import "github.com/gophertrain/modules/usefulinterfaces/includes/impl"

// CustomerService implements the CustomerGetter interface
type CustomerService struct {
}

// Get returns a customer
func (c *CustomerService) Get(id int) (*impl.Customer, error) {
	panic("not implemented")
}
