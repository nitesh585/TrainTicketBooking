FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/userService

COPY . /go/userService
RUN go get ./... 

RUN go build -o /go/userService

EXPOSE 8081

CMD [ "/go/userService/userService" ]
