FROM golang:1.12

WORKDIR /usr/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build .

EXPOSE 80

ENTRYPOINT ["./loaddriver-master"]