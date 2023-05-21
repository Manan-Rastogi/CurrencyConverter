package main

import (
	"context"
	"fmt"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/manan-rastogi/currencyConvertor/helpers"
	"github.com/manan-rastogi/currencyConvertor/services"
)

func (s *CurrencyServer) CurrencyConverterServerStreaming(inputs *currencyProto.CurrencyRequestList, stream currencyProto.Converter_CurrencyConverterServerStreamingServer) error {
	helpers.InfoLog("Server Streaming Started.")

	service := services.NewService()
	for _, input := range inputs.Requests {
		response, err := service.CurrencyConveterService(context.Background(), input)
		if err != nil{
			helpers.ErrorLog(fmt.Sprintf("Error in getting response for input - %v", input))
			// Not returning here as we wish to continue for all inputs
		}

		if err := stream.Send(response); err != nil {
			helpers.ErrorLog(fmt.Sprintf("Unable to stream responses back - %v from input - %v", err.Error(), input))
			return err
		} 
	}

	return nil
}
