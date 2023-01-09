FROM cgr.dev/chainguard/go:latest

WORKDIR /usr/code
COPY . .
RUN go mod tidy

ENTRYPOINT ["go run main.go"]


