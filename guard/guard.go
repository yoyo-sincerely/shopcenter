package guard

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
)

type Guardian struct {
	conn db.Connection
}

func New( /*...*/ ) *Guardian {
	return &Guardian{
		// ...
	}
}

func (g *Guardian) Update( /*...*/ ) {
	// ...
}
