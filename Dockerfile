FROM golang:1.21-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/api ./cmd/api 

CMD ["/app/api"]

