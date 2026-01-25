FROM golang:1.25-alpine

RUN apk add --no-cache curl

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /learning-go ./cmd/api/main.go

# Exponemos el puerto de Fiber
EXPOSE 3000

# HEALTHCHECK: Usa tu endpoint de health para avisar a Docker si el proceso está vivo
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:3000/ping || exit 1

# Ejecución
CMD ["/learning-go"]