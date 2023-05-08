package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ControllerTransaction interface {
	GetTransaction(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	AddTransaction(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	UpdateTransaction(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
	DeleteTransaction(writer http.ResponseWriter, request *http.Request, Params httprouter.Params)
}
