FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN apk add git

RUN go mod download

COPY . ./

RUN go build -o easyfitanalyzer

EXPOSE 8080

CMD ["./easyfitanalyzer"]
