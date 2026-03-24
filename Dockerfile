FROM golang:1.26

WORKDIR /books

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main cmd/api/main.go

EXPOSE 8080

CMD [ "/books/main" ]