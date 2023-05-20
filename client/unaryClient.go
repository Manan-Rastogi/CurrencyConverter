package main

import (
	"context"

	"time"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/manan-rastogi/currencyConvertor/helpers"
)

func unaryCurrencyConverter(client currencyProto.ConverterClient, input *currencyProto.CurrencyRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	welcome, err := client.CurrencyConverter(ctx, input)

	if err != nil {
		helpers.ErrorLogF("Got Error from unary server: " + err.Error())
	}

	helpers.ValueLog("Unary Response", welcome)
}
