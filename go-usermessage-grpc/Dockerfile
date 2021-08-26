FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /go_usermessage_grpc
RUN mkdir -p /go_usermessage_grpc/usermsg 

WORKDIR /go_usermessage_grpc

COPY ./usermsg /go_usermessage_grpc
COPY ./usermsg_server/usermsg_server.go /go_usermessage_grpc

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o go_usermessage_grpc .

EXPOSE 50515

CMD ./go_usermessage_grpc