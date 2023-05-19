package main

import (
	"context"

	"github.com/fatih/color"
	"github.com/manan-rastogi/currencyConvertor/currencyProto"
)

func (s *CurrencyServer) CurrencyConverter(context.Context, *currencyProto.NoParam) (*currencyProto.Welcome, error) {
	color.Green("Client Request Received.")
	return &currencyProto.Welcome{
		Welcome:  "Welcome to Currency Converter.",
	}, nil
}
