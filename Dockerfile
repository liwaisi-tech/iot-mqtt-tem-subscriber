FROM golang:1.22.4-alpine

# Add Maintainer Info
LABEL maintainer="Braejan Arias <ing.brayanarias@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Set Environment Variables
ENV PG_DATABASE_HOST=${PG_DATABASE_HOST} \
    PG_DATABASE_USER=${PG_DATABASE_USER} \
    PG_DATABASE_PASSWORD=${PG_DATABASE_PASSWORD} \
    PG_DATABASE_NAME=${PG_DATABASE_NAME} \
    PG_DATABASE_PORT=${PG_DATABASE_PORT} \
    GORM_MAX_IDLE_CONNS=${GORM_MAX_IDLE_CONNS} \
    GORM_MAX_OPEN_CONNS=${GORM_MAX_OPEN_CONNS} \
    GORM_MAX_LIFE_TIME=${GORM_MAX_LIFE_TIME} \
    ENV=${ENV} \
    MQTT_BROKER=${MQTT_BROKER} \
    MQTT_CLIENT_ID=${MQTT_CLIENT_ID} \
    MQTT_USERNAME=${MQTT_USERNAME} \
    MQTT_PASSWORD=${MQTT_PASSWORD} \
    MQTT_TOPIC=${MQTT_TOPIC}

# Build the Go app
RUN go build -o main ./cmd/cli/main.go

# Command to run the executable
CMD ["./main"]