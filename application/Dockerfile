#
FROM golang:1.18 as builder

WORKDIR /build
COPY utils /build/utils/
COPY main.go /build/
COPY go.mod /build/
COPY go.sum /build/

RUN go build -o /build/blog
COPY static /build/static

CMD [ "/build/blog" ]
