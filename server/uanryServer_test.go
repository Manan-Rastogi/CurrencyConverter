package main

import (
	"context"
	"log"
	"testing"

	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/stretchr/testify/require"
)

var server = NewServer()

func TestCurrencyConverter(t *testing.T) {
	log.Println("Testing for Unary Currency Converter.")
	ctx := context.Background()
	input := currencyProto.CurrencyRequest{
		From: "INR",
		To: "USD",
		Amount: 1200,
	}

	unaryResponse, err := server.CurrencyConverter(ctx, &input)
	require.NoError(t, err)
	require.NotZero(t, unaryResponse)
}
