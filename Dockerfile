FROM golang:1.21

RUN mkdir -p /go/src/testing-nextalent
WORKDIR /go/src/testing-nextalent
COPY . .

RUN go mod download

RUN go build -o /main main.go

EXPOSE 8080

ENTRYPOINT ["/main"]