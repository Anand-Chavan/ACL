package user

import (
	"context"
	"database/sql"
	//"fmt"
	"github.com/pucsd2020-pp/ACL/driver"
	"github.com/pucsd2020-pp/ACL/model"
)

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *userRepository {
	//fmt.Println("repository->user->user.go->NewUserRepository")
	return &userRepository{conn: conn}
}

func (user *userRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	//fmt.Println("repository->user->user.go->GetById")
	obj := new(model.User)
	return driver.GetById(user.conn, obj, id)
}

func (user *userRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	//usr := obj.(*model.User)
	//result, err := driver.Create(user.conn, usr)
	//fmt.Println("repository->user->user.go->Create")
	usr := obj.(model.User)
	result, err := driver.Create(user.conn,&usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (user *userRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	//fmt.Println("repository->user->user.go->Update")
	usr := obj.(model.User)
	err := driver.UpdateById(user.conn, &usr)
	return obj, err
}

func (user *userRepository) Delete(cntx context.Context, id int64) error {
	//fmt.Println("repository->user->user.go->Delete")
	obj := model.User{Id: id}
	return driver.SoftDeleteById(user.conn, &obj, id)
}

func (user *userRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	//fmt.Println("repository->user->user.go->GetAll")
	obj := &model.User{}
	return driver.GetAll(user.conn, obj, 0, 0)
}
