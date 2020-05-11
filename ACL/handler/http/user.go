package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"

	"github.com/go-chi/chi"

	"github.com/pucsd2020-pp/ACL/handler"
	"github.com/pucsd2020-pp/ACL/model"
	"github.com/pucsd2020-pp/ACL/repository"
	"github.com/pucsd2020-pp/ACL/repository/user"
)

type User struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewUserHandler(conn *sql.DB) *User {
	//fmt.Println("handler->http ->  user.go->NewUserHandler")
	return &User{
		repo: user.NewUserRepository(conn),
	}
}

func (user *User) GetHTTPHandler() []*handler.HTTPHandler {
	//fmt.Println("handler->http ->  user.go->GetHTTPHandler")
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user/{id}", Func: user.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "user", Func: user.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "user/{id}", Func: user.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "user/{id}", Func: user.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user", Func: user.GetAll},
	}
}

func (user *User) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
	for {
		if nil != err {
			break
		}

		usr, err = user.repo.GetByID(r.Context(), id)
		break
	}

	//fmt.Println("handler->http ->  user.go->GetByID")

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *User) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = user.repo.Create(r.Context(), usr)
		break
	}

	fmt.Println("handler->http ->  user.go->create")

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}
