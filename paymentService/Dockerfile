FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/paymentService

COPY . /go/paymentService
RUN go get ./... 

RUN go build -o /go/paymentService

EXPOSE 8083

CMD [ "/go/paymentService/paymentService" ]
