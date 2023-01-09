FROM golang:1.18-alpine as build

WORKDIR /logchain

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o logchain main.go

FROM alpine 

WORKDIR /logchain

COPY --from=build /logchain/logchain .

COPY config.yaml .

EXPOSE 9090

ENTRYPOINT [ "/logchain/logchain","start" ]