package repository

import (
	"context"
)

type IRepository interface {
	GetByID(context.Context, string) (interface{}, error)
	Create(context.Context, interface{}) (interface{}, error)
	Authentication(context.Context, interface{}) (interface{}, error)
	Logout(context.Context, interface{}) (interface{}, error)
	Update(context.Context, interface{}) (interface{}, error)
	Delete(context.Context, string) error
	GetAll(context.Context) ([]interface{}, error)
	SetAll(context.Context) ([]interface{}, error)
	GetByuserId(context.Context, string) (interface{}, error)
}

type Repository struct {
}

func (repo *Repository) GetByID(cntx context.Context, id int64) (obj interface{}, err error) {
	return
}

func (repo *Repository) Create(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}
func (repo *Repository) Authentication(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}
func (repo *Repository) Logout(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

func (repo *Repository) Update(cntx context.Context, obj interface{}) (uobj interface{}, err error) {
	return
}

func (repo *Repository) Delete(cntx context.Context, id int64) (deleted bool, err error) {
	return
}

func (repo *Repository) GetAll(cntx context.Context) (obj []interface{}, err error) {
	return
}
func (repo *Repository) SetAll(cntx context.Context) (obj []interface{}, err error) {
	return
}

func (repo *Repository) GetByuserId(cntx context.Context, id string) (obj interface{}, err error) {
	return
}
