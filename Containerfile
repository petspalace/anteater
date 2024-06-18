FROM docker.io/library/golang:1.22-alpine AS build
MAINTAINER Simon de Vlieger <supakeen@redhat.com>

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./cmd ./cmd
COPY *.go ./

RUN go build -o /anteater ./cmd/anteater

FROM docker.io/library/alpine:latest
MAINTAINER Simon de Vlieger <supakeen@redhat.com>

COPY --from=build /anteater /anteater

CMD ["/anteater"]
