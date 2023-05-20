package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/manan-rastogi/currencyConvertor/helpers"
	"github.com/manan-rastogi/currencyConvertor/services"
)

func (s *CurrencyServer) CurrencyConverterClientStreaming(stream currencyProto.Converter_CurrencyConverterClientStreamingServer) (err error) {
	helpers.InfoLog("Got Client Streaming Requests.")
	var responses []*currencyProto.CurrencyResponse
	service := services.NewService()
	
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			helpers.InfoLog("End of Incoming Stream -- " + err.Error())
			return stream.SendAndClose(&currencyProto.CurrencyResponseList{
				Responses: responses,
			})
		}
		if err != nil {
			helpers.ErrorLog("Unexpected Error occured -- " + err.Error())
			return err
		}
	
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()
		response, err := service.CurrencyConveterService(ctx, req)
		if err != nil{
			helpers.ErrorLog(fmt.Sprintf("error - %v in getting response for - %v", err.Error(), req))
			// Not returning here as we want to continue for rest of the requests
		}

		responses = append(responses, response)
	}	

}
