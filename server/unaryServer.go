package main

import (
	"context"

	
	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/manan-rastogi/currencyConvertor/helpers"
	"github.com/manan-rastogi/currencyConvertor/services"
	
)

func (s *CurrencyServer) CurrencyConverter(ctx context.Context, input *currencyProto.CurrencyRequest) (*currencyProto.CurrencyResponse, error) {
	helpers.InfoLog("Unary Request Received")
	
	service := services.NewService()

	return service.CurrencyConveterService(context.Background(), input)
}
