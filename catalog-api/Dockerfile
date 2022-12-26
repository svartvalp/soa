FROM golang:1.19

WORKDIR $GOPATH/src/github.com/soa/catalog-api
COPY ./ ./
RUN go build -o ./bin/main ./cmd/catalog-api/main.go
CMD ["./bin/main"]
