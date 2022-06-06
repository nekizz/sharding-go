FROM golang:1.18

WORKDIR /go/src/sharding

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go/src/sharding

EXPOSE 3000

CMD [ "/sharding" ]

