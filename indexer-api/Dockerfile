FROM golang:1.19

WORKDIR $GOPATH/src/github.com/soa/indexer-api
COPY ./ ./
RUN go build -o ./bin/main ./cmd/indexer-api/main.go
CMD ["./bin/main"]
