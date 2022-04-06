FROM golang:1.17-alpine

RUN mkdir /app

WORKDIR /app

ADD ./ /app

RUN go mod download

RUN go build -o main .

CMD [ "/app/main" ]