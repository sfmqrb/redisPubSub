FROM golang:1.16-alpine
WORKDIR /subscriber
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /go-build
CMD [ "/go-build" ]
