FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o temperature-api .

EXPOSE 8081

CMD ["./temperature-api"]