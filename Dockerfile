FROM golang:1.14-alpine

RUN apk update

RUN apk add --no-cache \
    gtk+3.0 \
    gtk+3.0-dev \
    glib \
    glib-dev \
    cairo \
    cairo-dev \
    gdk-pixbuf \
    gdk-pixbuf-dev

WORKDIR /go/src/app

ENV GO111MODULE=on
#ENV GOPROXY=off

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD . .

RUN go build -o ipfs-cloud

ENTRYPOINT [ "./ipfs-cloud" ]
