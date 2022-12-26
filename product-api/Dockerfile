FROM golang:1.19

WORKDIR $GOPATH/src/github.com/soa/product-api
COPY ./ ./
RUN go build -o ./bin/main ./cmd/product-api/main.go
CMD ["./bin/main"]
