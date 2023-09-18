# Usar una imagen base de Go
FROM golang:1.20

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /go/src/github.com/Nachofra/final-esp-backend-3

# Copiar el archivo go.mod y go.sum para descargar las dependencias
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar todo el código fuente desde la ubicación relativa "cmd/api/" al directorio de trabajo actual
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/
COPY ./pkg/ ./pkg/

# Ejecutar go mod tidy
RUN go mod tidy

# Compilar la aplicación
RUN go build -o main ./cmd/api/

# Exponer el puerto en el que la aplicación escuchará
EXPOSE 8080

# Comando para ejecutar la aplicación al iniciar el contenedor
CMD ["./main"]