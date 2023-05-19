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
	noParam := currencyProto.NoParam{}

	unaryResponse, err := server.CurrencyConverter(ctx, &noParam)
	require.NoError(t, err)
	require.NotZero(t, unaryResponse)
	require.Equal(t, "Welcome to Currency Converter.", unaryResponse.Welcome)
}
