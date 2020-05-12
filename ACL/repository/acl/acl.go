package acl

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type aclRepository struct {
	conn *sql.DB
}

func NewAclRepository(conn *sql.DB) *aclRepository {
	return &aclRepository{conn: conn}
}

func (acl *aclRepository) GetByID(cntx context.Context, id string) (interface{}, error) {
	obj := new(model.Acl)
	return driver.GetById(acl.conn, obj, id)
}

func (acl *aclRepository) GetByuserId(cntx context.Context, id string) (interface{}, error) {
	obj := new(model.Acl)
	return driver.GetByuserId(acl.conn, obj, id)
}

func (acl *aclRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.Acl)
	result, err := driver.Create(acl.conn, &usr)
	if nil != err {
		return 0, err
	}

	// id, _ := result.LastInsertId()
	// usr.Id = id
	return result.RowsAffected, nil
}
func (acl *aclRepository) Authentication(cntx context.Context, obj interface{}) (interface{}, error) {
	auth := obj.(model.Auth)
	usr1, err := driver.Authentication(acl.conn, &auth)
	if nil != err {
		return 0, err
	}
	return usr1, nil
}
func (acl *aclRepository) Logout(cntx context.Context, obj interface{}) (interface{}, error) {
	auth := obj.(model.Key)
	usr1, err := driver.Logout(acl.conn, &auth)
	if nil != err {
		return 0, err
	}
	return usr1, nil
}

func (acl *aclRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.Acl)
	err := driver.UpdateById(acl.conn, &usr)
	return obj, err
}

func (acl *aclRepository) Delete(cntx context.Context, id string) error {
	obj := &model.Acl{UserId: id}
	return driver.SoftDeleteById(acl.conn, obj, id)
}

func (acl *aclRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.Acl{}
	return driver.GetAll(acl.conn, obj, 0, 0)
}
func (acl *aclRepository) SetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.Acl{}
	return driver.GetAll(acl.conn, obj, 0, 0)
}
