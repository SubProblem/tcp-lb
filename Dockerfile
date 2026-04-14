FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o tcplb .


FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /app/tcplb .
EXPOSE 8080
ENTRYPOINT ["./tcplb"]