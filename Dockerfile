FROM golang

WORKDIR /root
ENV HOME=/root

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY database.go database.go

RUN go build -o main .

CMD ["./main"]