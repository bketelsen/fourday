package pool

// START OMIT
// Pool holds Things
type Pool struct {
	pool chan *Thing
}

// END OMIT

// STARTNEW OMIT
// NewPool creates a new pool of Things
func NewPool(max int) *Pool {
	return &Pool{
		pool: make(chan *Thing, max),
	}
}

// ENDNEW OMIT

// STARTBORROW OMIT
// Borrow a Thing from the pool.
func (p *Pool) Borrow() *Thing {
	var t *Thing
	select {
	case t = <-p.pool:
	default:
		t = newThing()
	}
	return t
}

// ENDBORROW OMIT

// STARTRETURN OMIT
// Return returns a Thing to the pool.
func (p *Pool) Return(t *Thing) {
	select {
	case p.pool <- t:
	default:
	}
}

// ENDRETURN OMIT
