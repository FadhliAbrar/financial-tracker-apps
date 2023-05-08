package helper

import (
	x "github.com/FadhliAbrar/financial-tracker-apps/Model/Domain"
	web "github.com/FadhliAbrar/financial-tracker-apps/Model/Web"
)

//x is refering to domain model. not the category struct

func ToCategoryWebResponse(domain x.Category) web.CategoryResponses {
	str, err := x.ToFlowString(domain.Type)
	PanicIfError(err)
	return web.CategoryResponses{
		Id:   domain.Id,
		Name: domain.Name,
		Type: str,
	}
}

func ToCategoryWebResponses(domain []x.Category) []web.CategoryResponses {
	responseCategories := []web.CategoryResponses{}
	for _, v := range domain {
		c := ToCategoryWebResponse(v)
		responseCategories = append(responseCategories, c)
	}
	return responseCategories
}

func ToTransactionResponse(domain x.Transaction) web.TransactionResponse {
	str, err := x.ToFlowString(domain.Type)
	PanicIfError(err)
	return web.TransactionResponse{
		Id:          domain.Id,
		Type:        str,
		Amount:      domain.Amount,
		Category_id: domain.Category_id,
		Date:        domain.Date,
	}
}

func ToTransactionResponses(domain []x.Transaction) []web.TransactionResponse {
	responseTransaction := []web.TransactionResponse{}
	for _, v := range domain {
		c := ToTransactionResponse(v)
		responseTransaction = append(responseTransaction, c)
	}
	return responseTransaction
}

func ToUserResponse(domain x.User) web.UserResponse {
	return web.UserResponse{
		Id:       domain.Id,
		Username: domain.Username,
		Email:    domain.Email,
	}
}
