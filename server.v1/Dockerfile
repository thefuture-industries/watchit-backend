FROM golang:1.23.0

WORKDIR /build_app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o build_app

# EXPOSE 8080

CMD ["/build_app/build_app"]
