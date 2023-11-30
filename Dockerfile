FROM golang:alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
EXPOSE 8080

RUN go build cmd/app/main.go

CMD ["./main"]