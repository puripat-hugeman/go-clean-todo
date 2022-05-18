FROM golang:1.18

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o /app/out ./cmd/main.go
 
EXPOSE 8080

CMD ["/app/out"]