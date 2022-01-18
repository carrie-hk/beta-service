#build stage
FROM golang:alpine as builder

# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0

#create the work directory for the app in the container
WORKDIR /app

#copy the mod and sum files into the container
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

#creates executable main file
RUN go build ./app/main.go

#exposes port 5000 to listen on
EXPOSE 5004
CMD ["./main"]
