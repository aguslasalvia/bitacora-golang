# Etapa 1: Builder
FROM golang:1.24-bullseye AS builder

# Directorio de trabajo
WORKDIR /app

# Instalar librerías necesarias para CGO (SQLite)
RUN apt-get update && \
    apt-get install -y build-essential libsqlite3-dev git && \
    rm -rf /var/lib/apt/lists/*

# Copiar módulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo el proyecto
COPY . .

# Compilar binario y dejar scripts/ y templates/ como hermanos
RUN mkdir -p bin && \
    go build -o bin/app ./cmd/bitacora && \
    cp -r ./scripts bin/ && \
    cp -r ./templates bin/

# Etapa 2: Imagen final mínima
FROM debian:11-slim

# Instalar librerías de runtime necesarias para CGO
RUN apt-get update && \
    apt-get install -y libsqlite3-0 ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copiar binario junto con scripts/ y templates/ (paths relativos al cwd)
COPY --from=builder /app/bin/ ./

# Definir comando por defecto
CMD ["./app"]
