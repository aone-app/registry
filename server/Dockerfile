FROM golang:1.25.1-alpine AS builder
WORKDIR /app
# COPY /src/go.mod /src/go.sum ./
COPY /src/ ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /src .

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /src /src
EXPOSE 8080
ENTRYPOINT ["/src"]