FROM golang:1.26 AS builder 

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /books
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o main cmd/api/main.go


FROM alpine:3.21 AS final

WORKDIR /books

ENV HOST=0.0.0.0
ENV PORT=8080
ENV BLUEPRINT_DB_URL="/books/prod.db"
ENV APP_ENV=prod

COPY --from=builder /books/main .
COPY --from=builder /books/templates ./templates
COPY --from=builder /books/prod.db .


EXPOSE 80

ENTRYPOINT  [ "/books/main" ]