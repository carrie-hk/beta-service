#build stage
FROM golang:alpine as builder


#create the work directory for the app in the container
WORKDIR /app

#copy the mod and sum files into the container
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

#creates executable main file
RUN go build ./app/main.go

#exposes port 443 to listen on
EXPOSE 443
CMD ["./main"]
