package main

import (
	"context"
	"io"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/manan-rastogi/currencyConvertor/helpers"
)

func callCurrencyConverterServerStreaming(client currencyProto.ConverterClient, inputs *currencyProto.CurrencyRequestList) {
	helpers.InfoLog("Server Streaming Strated.")

	stream, err := client.CurrencyConverterServerStreaming(context.Background(), inputs)
	if err != nil{
		helpers.ErrorLogF("Error sending request to server - " + err.Error())
	}

	for{
		response, err := stream.Recv()
		if err == io.EOF{
			helpers.InfoLog("End of Stream.")
			break
		}
		if err != nil{
			helpers.ErrorLogF("Error while streaming - " + err.Error())
		}
		
		helpers.ValueLog("Response", response)
	}

	helpers.InfoLog("Streaming has finished.")
}
