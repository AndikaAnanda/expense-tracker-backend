# Build Stage
FROM golang:1.25-alpine AS builder
WORKDIR /app

# Install git for fetching dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o expense-tracker-backend .

# Final Stage
FROM alpine:latest

WORKDIR /app

# Install tzdata for timezone support
RUN apk add --no-cache tzdata

COPY --from=builder /app/expense-tracker-backend .
COPY --from=builder /app/.env .

EXPOSE 8081

CMD ["./expense-tracker-backend"]