FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY /cmd ./cmd
COPY /pkg ./pkg
RUN go build -o ./cyber-meower-query-service ./cmd

CMD [ "./cyber-meower-query-service" ]