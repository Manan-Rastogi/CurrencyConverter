package main

import (
	"context"
	"io"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/manan-rastogi/currencyConvertor/helpers"
)

func callCurrencyConverterBidirectionalStreaming(client currencyProto.ConverterClient, inputs *currencyProto.CurrencyRequestList) {
	helpers.InfoLog("Bidirectional Streaming strated.")

	stream, err := client.CurrencyConverterBidirectionalStreaming(context.Background())
	if err != nil {
		helpers.ErrorLogF("Error sending request to server - " + err.Error())
	}


	waitChannel := make(chan struct{})
	// Lets Use routine to receive response..

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				helpers.InfoLog("End of Stream.")
				break
			}
			if err != nil {
				helpers.ErrorLogF("Error while streaming - " + err.Error())
			}
	
			helpers.ValueLog("Response", response)
		}
		close(waitChannel)
	}()



	for _, input:= range inputs.Requests{
		if err := stream.Send(input); err != nil {
			helpers.ErrorLogF("Error sending inputs via client Stream - " + err.Error())
		}
		helpers.ValueLog("Request Send with", input)
	}

	stream.CloseSend()
	<-waitChannel

	helpers.InfoLog("Bidirectional streaming finished")
}
