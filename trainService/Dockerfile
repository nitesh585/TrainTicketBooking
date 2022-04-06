FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/trainService

COPY . /go/trainService
RUN go get ./... 

RUN go build -o /go/trainService

EXPOSE 8081

CMD [ "/go/trainService/trainService" ]
