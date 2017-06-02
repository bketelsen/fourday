package transport

import (
	"net"

	"github.com/gophertrain/material/usefulinterfaces/solutions/inventory"
)

type InventoryTransporter interface {
	inventory.Service
	Serve(net.Listener) error
}
