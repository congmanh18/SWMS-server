# Use the official Golang image
FROM golang:1.22.0-alpine

# Install git and other dependencies
RUN apk update && apk add --no-cache git

# Set environment variables
ENV GO111MODULE=on
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

# Create necessary directories and set permissions
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Set the working directory inside the container
WORKDIR $GOPATH/src/smart-waste-system

# Copy everything from the current directory to the working directory in the container
COPY . .

RUN go mod init smart-waste-system

# Add missing dependencies
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-playground/validator/v10
RUN go get github.com/gofiber/fiber/v2
RUN go get github.com/google/uuid
RUN go get github.com/joho/godotenv
RUN go get golang.org/x/crypto/bcrypt
RUN go get gorm.io/driver/postgres
RUN go get gorm.io/gorm
RUN go get github.com/go-playground/validator
RUN go get github.com/gofiber/websocket/v2


# Install necessary dependencies
RUN go get -d ./...

# Build the Go app
RUN GOOS=linux go build -o app .

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./app"]


