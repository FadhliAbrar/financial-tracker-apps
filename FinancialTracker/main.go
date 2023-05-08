package main

import (
	"net/http"

	controller "github.com/FadhliAbrar/financial-tracker-apps/Controller"
	database "github.com/FadhliAbrar/financial-tracker-apps/Database"
	helper "github.com/FadhliAbrar/financial-tracker-apps/Helper"
	repository "github.com/FadhliAbrar/financial-tracker-apps/Repository"
	services "github.com/FadhliAbrar/financial-tracker-apps/Services"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := database.GetDatabase()

	router := httprouter.New()
	transactionRepository := repository.NewTransactionRepository()
	transactionService := services.NewTransactionService(db, transactionRepository)
	controllers := controller.NewTransactionController(transactionService)

	router.GET("/api/transactions/", controllers.GetTransaction)
	router.POST("/api/transactions/", controllers.AddTransaction)
	router.PUT("/api/transactions/", controllers.UpdateTransaction)
	router.DELETE("/api/transactions/:transaction_id", controllers.DeleteTransaction)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
