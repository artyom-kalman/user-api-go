FROM golang:1.24

WORKDIR /app

COPY ../go.mod .
COPY ../go.sum .
RUN go mod download

COPY ../. .

RUN go build -o main cmd/app/main.go

EXPOSE 3030

CMD ["./main"]
