run-users-services:
	go run services/users/cmd/main.go

run-books-services:
	go run services/books/cmd/main.go

run-generate-proto:
	protoc --go_out=. --go-grpc_out=. ./proto/books.proto