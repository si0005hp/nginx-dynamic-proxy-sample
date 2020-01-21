FROM golang:1.11-stretch as build
WORKDIR /go/src
COPY ap/http.go .
RUN go build -o http http.go

FROM ubuntu:18.04
WORKDIR /go/src
COPY --from=build /go/src/http .

CMD ["./http"]