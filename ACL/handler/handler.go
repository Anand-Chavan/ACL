package handler

import (
	"encoding/json"
	"net/http"
	// "fmt"
)

type IHTTPHandler interface {
	GetHTTPHandler() []*HTTPHandler
	GetByID(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}

type HTTPHandler struct {
	Authenticated bool
	Method        string
	Path          string
	Func          func(http.ResponseWriter, *http.Request)
}

type response struct {
	Status  int         `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (hdlr *HTTPHandler) GetHTTPHandler() []HTTPHandler {
	//fmt.Println("handler->hnadler.go->GetHTTPHandler")
	return []HTTPHandler{}
}

func (hdlr *HTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("handler->hnadler.go->GetByID")
	return
}

func (hdlr *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("handler->hnadler.go->Create")
	return
}

func (hdlr *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("handler->hnadler.go->Update")
	return
}

func (hdlr *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("handler->hnadler.go->Delete")
	return
}

func (hdlr *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("handler->hnadler.go->GetAll")
	return
}

func WriteJSONResponse(w http.ResponseWriter,
	r *http.Request,
	payload interface{},
	code int,
	err error) {
	resp := &response{
		Status: code,
		Data:   payload,
	}

	if nil != err {
		resp.Message = err.Error()
	}

	response, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	//fmt.Println("handler->hnadler.go->WriteJSONResponse")
	return
}
