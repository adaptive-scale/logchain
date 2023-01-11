generate:
	 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/logchain/*.proto

build:
	rm -f logchain
	go build -o logchain ./main.go

docker:
	docker build . -t registry.digitalocean.com/adaptive-at-scale/logchain:latest
	docker push registry.digitalocean.com/adaptive-at-scale/logchain:latest
update:
	go mod tidy
