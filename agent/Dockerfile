FROM golang:latest as builder

WORKDIR /app

COPY ./agent/go.mod ./agent/go.sum ./

RUN go mod download

COPY ./agent .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8000 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]