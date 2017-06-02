package mocks

import (
	"errors"
	"sync"

	"github.com/gophertrain/material/usefulinterfaces/solutions/inventory"
)

var _ inventory.OrderStorage = (*OrderMock)(nil)

type OrderMock struct {
	mu     sync.Mutex
	orders map[int]*inventory.Order
}

func NewOrderMock() *OrderMock {
	return &OrderMock{
		orders: make(map[int]*inventory.Order),
	}
}

func (s *OrderMock) Get(id int) (*inventory.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o, ok := s.orders[id]
	if !ok {
		return nil, errors.New("Record not found")
	}
	return o, nil
}

func (s *OrderMock) Create(o inventory.Order) (*inventory.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orders[o.ID] = &o
	return s.orders[o.ID], nil
}
