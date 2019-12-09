FROM golang:1.12

RUN apt-get update -y && apt-get install openjdk-11-jre-headless -y

WORKDIR /usr/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build .

EXPOSE 80

ENTRYPOINT ["./loaddriver-master"]