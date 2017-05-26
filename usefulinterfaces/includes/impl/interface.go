package impl

type Customer struct {
	ID   int
	Name string
}

type CustomerGetter interface {
	Get(id int) (*Customer, error)
}
