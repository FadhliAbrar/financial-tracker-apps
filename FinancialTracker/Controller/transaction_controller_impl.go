package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	helper "github.com/FadhliAbrar/financial-tracker-apps/Helper"
	web "github.com/FadhliAbrar/financial-tracker-apps/Model/Web"
	services "github.com/FadhliAbrar/financial-tracker-apps/Services"
	"github.com/julienschmidt/httprouter"
)

type transactionControllerImpl struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) ControllerTransaction {
	return &transactionControllerImpl{
		service: service,
	}
}

func (controller *transactionControllerImpl) GetTransaction(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	transactionResponse := controller.service.FindAllTransaction(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *transactionControllerImpl) AddTransaction(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	decode := json.NewDecoder(request.Body)
	transactionRequest := web.TransactionRequest{}
	err := decode.Decode(&transactionRequest)
	helper.PanicIfError(err)

	transactionResponse := controller.service.CreateTransaction(request.Context(), transactionRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *transactionControllerImpl) UpdateTransaction(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	decode := json.NewDecoder(request.Body)
	transactionRequest := web.TransactionRequest{}
	err := decode.Decode(&transactionRequest)
	helper.PanicIfError(err)

	transactionResponse := controller.service.UpdateTransaction(request.Context(), transactionRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transactionResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
func (controller *transactionControllerImpl) DeleteTransaction(writer http.ResponseWriter, request *http.Request, Params httprouter.Params) {
	transaction_id, err := strconv.Atoi(Params.ByName("transaction_id"))
	helper.PanicIfError(err)

	controller.service.DeleteTransactionById(request.Context(), transaction_id)
	webResponse := web.WebResponseDelete{
		Code:   200,
		Status: "OK",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
