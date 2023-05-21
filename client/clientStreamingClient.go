package main

import (
	"context"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/manan-rastogi/currencyConvertor/helpers"
)

func callCurrencyConverterClientStreaming(client currencyProto.ConverterClient, inputs *currencyProto.CurrencyRequestList) {
	helpers.InfoLog("Client started streaming.")

	stream, err := client.CurrencyConverterClientStreaming(context.Background())
	if err != nil {
		helpers.ErrorLogF("Error generating client Stream - " + err.Error())
	}

	for _, input := range inputs.Requests {

		if err := stream.Send(input); err != nil {
			helpers.ErrorLogF("Error sending inputs via client Stream - " + err.Error())
		}
		helpers.ValueLog("Request Send with", input)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		helpers.ErrorLogF("Error closing client Stream - " + err.Error())
	}

	helpers.InfoLog("Client Streaming Finished")

	helpers.ValueLog("Response", response)
}
