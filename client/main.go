package main

import (
	"log"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err.Error())
	}

	defer conn.Close()

	client := currencyProto.NewConverterClient(conn)

	clientRequestStream := currencyProto.CurrencyRequestList{
		Requests: []*currencyProto.CurrencyRequest{
			{
				From:   "INR",
				To:     "USD",
				Amount: 1200,
			},
			{
				From:   "EUR",
				To:     "CAD",
				Amount: 1200,
			},
			{
				From:   "AUD",
				To:     "CNY",
				Amount: 1200,
			},
			{
				From:   "BRL",
				To:     "JPY",
				Amount: 1200,
			},
		},
	}

	// unaryCurrencyConverter(client,  &currencyProto.CurrencyRequest{
	// 	To: "INR",
	// 	From: "AUD",
	// 	Amount: 15000,
	// })

	callCurrencyConverterClientStreaming(client, &clientRequestStream)

}
