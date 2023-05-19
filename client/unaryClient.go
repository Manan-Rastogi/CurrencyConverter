package main

import (
	"context"
	"log"

	"time"

	"github.com/fatih/color"
	"github.com/manan-rastogi/currencyConvertor/currencyProto"
)

func unaryCurrencyConverter(client currencyProto.ConverterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	
	defer cancel()

	welcome, err := client.CurrencyConverter(ctx, &currencyProto.NoParam{})
	if err != nil{
		errMsg := color.RedString("Got Error from unary server: %v", err.Error())
		log.Fatalf(errMsg)
	}

	color.Green("Unary Response: %v", welcome)
}

