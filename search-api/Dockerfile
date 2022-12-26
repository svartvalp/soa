FROM golang:1.19

WORKDIR $GOPATH/src/github.com/soa/search-api
COPY ./ ./
RUN go build -o ./bin/main ./cmd/search-api/main.go
CMD ["./bin/main"]
