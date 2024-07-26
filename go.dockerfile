
FROM golang:latest as build

WORKDIR /api

COPY . .

# Download and install the dependencies:
RUN go get -d -v ./...

# Build the go app
RUN go build -o api cmd/api/main.go 


EXPOSE 8000

CMD ["./api"]