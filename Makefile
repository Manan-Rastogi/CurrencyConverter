gen: 
	protoc --go_out=. --go-grpc_out=. proto/*.proto

test_servers:
	go test -cover -v server/*.go

.PHONY: gen test_servers