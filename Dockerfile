FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

RUN go build -o main main.go

CMD [ "./main" ]
