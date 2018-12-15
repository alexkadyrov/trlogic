FROM golang:1.11-alpine
EXPOSE 8080
#RUN go build /api/main.go
#COPY * /go/src/
ADD / /go/
WORKDIR /go/src
RUN go build api
