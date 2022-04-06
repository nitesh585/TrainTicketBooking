FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/loggingService

COPY . /go/loggingService
RUN go get ./... 

RUN go build -o /go/loggingService

EXPOSE 8085

CMD [ "/go/loggingService/loggingService" ]
