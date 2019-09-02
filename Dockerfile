FROM golang:1.12.5-alpine

WORKDIR /go/src/nats-stream-cqrs-frontend
COPY . .
 
RUN go install

EXPOSE 8080

CMD [ "nats-stream-cqrs-frontend" ]