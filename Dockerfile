FROM golang:1.17-latest AS build

ENV PATH=*/go/bin:${PATH}
ENV CGO_ENABLED=0
ENV GO1111MODULE=on

# Create a directory for the project and use it as working directory
RUN mkdir /go/src/fitpiece/
WORKDIR /go/src/fitpiece

# Copy all the Code and stuff to compile everything
COPY . .

WORKDIR /go/src/fitpiece/main_server
# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o fitpiece_app .
#
#
# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

WORKDIR /app

# Copy the generated binary from builder image to execution image
COPY --from=builder /go/src/fitpiece/main_server/fitpiece_app ./fitpiece_app

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
ENTRYPOINT ["./fitpiece_app", "serve", "--url", "0.0.0.0", "--port", "8080", "--prod=true"]

CMD ["./fitpiece_app"]
