FROM golang:1.11
EXPOSE 8080
# create a working directory
WORKDIR /go/src/app
# run main.go
CMD ["go", "run", "src/main.go"]