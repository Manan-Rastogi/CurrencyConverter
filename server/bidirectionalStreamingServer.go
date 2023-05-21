package main

import (
	"context"
	"fmt"
	"io"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/manan-rastogi/currencyConvertor/helpers"
	"github.com/manan-rastogi/currencyConvertor/services"
)

func (s *CurrencyServer) CurrencyConverterBidirectionalStreaming(stream currencyProto.Converter_CurrencyConverterBidirectionalStreamingServer) error {
	helpers.InfoLog("Bidirectional Stream Started.")

	service := services.NewService()

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			helpers.InfoLog("Client Stream Ended.")
			return nil
		}
		if err != nil {
			helpers.ErrorLog("Error while streaming Client - " + err.Error())
			return err
		}

		helpers.ValueLog("Request", request)

		response, err := service.CurrencyConveterService(context.Background(), request)
		if err != nil {
			helpers.ErrorLog(fmt.Sprintf("Error in getting response for input - %v", request))
			// Not returning here as we wish to continue for all inputs
		}

		err = stream.Send(response)
		if err != nil {
			helpers.ErrorLog(fmt.Sprintf("Unable to stream responses back - %v from input - %v", err.Error(), request))
			return err
		}
	}

	
}
