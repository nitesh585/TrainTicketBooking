FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/invoiceService

COPY . /go/invoiceService
RUN go get ./... 

RUN go build -o /go/invoiceService

EXPOSE 8084

CMD [ "/go/invoiceService/invoiceService" ]
