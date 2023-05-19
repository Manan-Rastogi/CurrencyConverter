package main

import "github.com/manan-rastogi/currencyConvertor/currencyProto"

type CurrencyServer struct{
	currencyProto.ConverterServer
}

func NewServer() *CurrencyServer {
	return &CurrencyServer{
	}	
}


