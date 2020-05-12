package http

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pucsd2020-pp/rest-api/handler"
	"github.com/pucsd2020-pp/rest-api/model"
	"github.com/pucsd2020-pp/rest-api/repository"
	"github.com/pucsd2020-pp/rest-api/repository/acl"
)

type Acl struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewAclHandler(conn *sql.DB) *Acl {
	return &Acl{
		repo: acl.NewAclRepository(conn),
	}
}

func (acl *Acl) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "acl/{userId}", Func: acl.GetByuserId},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "acl", Func: acl.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "auth", Func: acl.Authentication},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "logout", Func: acl.Logout},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "acl/{userId}", Func: acl.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "acl/{userId}", Func: acl.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "acl", Func: acl.GetAll},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "acl1", Func: acl.SetAll},
		///write path =books/{id}
	}
}

func (acl *Acl) GetByuserId(w http.ResponseWriter, r *http.Request) {
	var usr interface{}

	userId := chi.URLParam(r, "userId")
	//id := chi.URLParam(r, "userId")
	//userId := chi.URLParam(r, "userId")
	// for {
	// 	// if nil != err {
	// 	// 	break
	// 	// }

	// 	usr, _ = acl.repo.GetByuserId(r.Context(), id)
	// 	break
	// }
	usr, _ = acl.repo.GetByuserId(r.Context(), userId)

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, errors.New("No Errors"))
}

func (acl *Acl) SetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := acl.repo.SetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}

func (acl *Acl) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}

	userId := chi.URLParam(r, "userId")
	usr, _ = acl.repo.GetByuserId(r.Context(), userId)

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, errors.New("No Errors"))
}
func (acl *Acl) Authentication(w http.ResponseWriter, r *http.Request) {
	var usr model.Auth
	var auth interface{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		auth, err = acl.repo.Authentication(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, auth, http.StatusOK, err)
}
func (acl *Acl) Logout(w http.ResponseWriter, r *http.Request) {
	var usr model.Key
	var auth interface{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		auth, err = acl.repo.Logout(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, auth, http.StatusOK, err)
}

func (acl *Acl) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.Acl
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = acl.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (acl *Acl) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	userId := chi.URLParam(r, "userId")
	usr := model.Acl{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.UserId = userId
		if nil != err {
			break
		}

		// set logged in Book id for tracking update
		//usr.UpdatedBy = 0

		iUsr, err = acl.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.Acl)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (acl *Acl) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	// id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	userId := chi.URLParam(r, "userId")
	for {
		// if nil != err {
		// 	break
		// }

		err := acl.repo.Delete(r.Context(), userId)
		if nil != err {
			break
		}
		payload = "Acl deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, errors.New("deleted success"))
}

func (acl *Acl) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := acl.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
