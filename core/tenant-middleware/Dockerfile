# Stage 1: Build
FROM golang:1.24-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app

# 1. Copiamos los archivos de gestión de dependencias de la RAÍZ
# En un monorepo, estos mandan sobre todo lo demás
COPY go.mod go.sum ./
RUN go mod download

# 2. Copiamos las librerías necesarias
COPY lib/ ./lib/

# 3. Copiamos el código del servicio específico
COPY core/tenant-middleware/ ./core/tenant-middleware/

# 4. Compilamos apuntando a la ruta del main desde la raíz
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/middleware ./core/tenant-middleware/cmd/main.go

# Stage 2: Runtime
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /app/middleware .
EXPOSE 8080
CMD ["./middleware"]