# Build
FROM golang:1.22 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ms-categories ./cmd/api

# Run
FROM golang:1.22
WORKDIR /app
COPY --from=builder /app/ms-categories .
EXPOSE 8080
CMD ["./ms-categories"]
