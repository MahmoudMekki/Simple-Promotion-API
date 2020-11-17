package internal

import (
	"database/sql"

	"github.com/promo/view"
)

type Plugin interface {
	CreatePromo(*sql.DB, *view.Promo) error
	GetPromo(int, *sql.DB) error
	GetAll(*sql.DB) (*[]view.Promo, error)
	UpdatePromo(int, *sql.DB, *view.Promo) error
	DeletePromo(int, *sql.DB) error
}

type prepare interface {
	Plug() Plugin
}

type UserInt struct{}

type INCInt struct{}

func (p *UserInt) Plug() Plugin {
	return &PromoUser{}
}

func (p *INCInt) Plug() Plugin {
	return &PromoInc{}
}
