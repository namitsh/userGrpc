FROM golang:1.17.6-alpine3.15 as builder

RUN apk add --no-cache git

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/server/main.go

RUN pwd

RUN echo $GOPATH

FROM alpine:3.15

RUN apk --no-cache add bash curl ca-certificates

RUN apk update

WORKDIR /bin

COPY --from=builder /usr/local/bin/app .

EXPOSE 50051

CMD ["app"]