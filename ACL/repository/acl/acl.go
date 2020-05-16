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
func (acl *aclRepository) GetGroupById(cntx context.Context, id string) (interface{}, error) {
	obj := new(model.UserAddToGroup)
	return driver.GetGroupById(acl.conn, obj, id)
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
func (acl *aclRepository) CreateFileFolder(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.CreateFileOrFolder)
	result, err := driver.CreateFileFolder(acl.conn, &usr)
	if nil != err {
		return 0, err
	}
	return result.RowsAffected, nil
}
func (acl *aclRepository) DeleteFileFolder(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.CreateFileOrFolder)
	result, err := driver.DeleteFileFolder(acl.conn, &usr)
	if nil != err {
		return 0, err
	}
	return result.RowsAffected, nil
}
func (acl *aclRepository) AddUserIntoGroup(cntx context.Context, obj interface{}) (interface{}, error) {
	grp := obj.(model.UserAddToGroup)
	result, err := driver.Create(acl.conn, &grp)
	if nil != err {
		return 0, err
	}

	// id, _ := result.LastInsertId()
	// usr.Id = id
	return result.RowsAffected, nil
}
func (acl *aclRepository) CreateGroup(cntx context.Context, obj interface{}) (interface{}, error) {
	grp := obj.(model.Groups)
	result, err := driver.CreateGroup(acl.conn, &grp)
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
func (acl *aclRepository) GetFilesFolder(cntx context.Context, obj interface{}) (interface{}, error) {
	// obj := &model.GetFilesFold{}
	// return driver.GetFilesFold(acl.conn, obj, 0, 0)
	auth := obj.(model.GetFilesFold)
	usr1, err := driver.GetFilesFold(acl.conn, &auth, 0, 0)
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
func (acl *aclRepository) ChangePermission(cntx context.Context, obj interface{}) (interface{}, error) {
	auth := obj.(model.ChangePermission)
	usr1, err := driver.ChangePermission(acl.conn, &auth)
	if nil != err {
		return 0, err
	}
	return usr1, nil
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
