FROM golang:latest 
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -u github.com/kobylyanskiy/dgraph-api/dgraph
RUN go get -u github.com/dgraph-io/dgo
RUN go build -o main .
CMD ["/app/main"]
