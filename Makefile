generate:
	 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/logchain/*.proto

build:
	rm logchain
	go build -o logchain ./main.go

update:
	go mod tidy