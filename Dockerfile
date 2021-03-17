## base image 
FROM golang:1.16-alpine3.13 

## add the html file to the container
ADD index.html /go/src/index.html

## add the web server code to the container
ADD webserver.go /go/src/webserver.go

## set the working directory
WORKDIR /go/src

RUN go mod init webserver
RUN go mod tidy 

## run go build to compile the binary
## executable of the web server
RUN go build -o main .

## get things started
CMD ["./main"]