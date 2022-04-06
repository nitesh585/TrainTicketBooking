FROM golang:1.17-alpine

RUN apk update && apk add --no-cache git build-base

WORKDIR /go/emailService

COPY . /go/emailService
RUN go get ./... 

RUN go build -o /go/emailService

EXPOSE 8082

CMD [ "/go/emailService/emailService" ]
