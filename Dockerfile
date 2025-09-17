FROM golang:1.25.1 as base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Final stage - Distroless image
FROM gcr.io/distroless/static-debian12

COPY --from=base  /app/main .

EXPOSE 8080

CMD ["./main"]