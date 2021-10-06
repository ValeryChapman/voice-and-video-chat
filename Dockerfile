FROM golang:1.16-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o chat-app ./cmd/main.go

CMD ["./chat-app"]