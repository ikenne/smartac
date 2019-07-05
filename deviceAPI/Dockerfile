ARG GO_VERSION=1.12
ARG ALPINE_VERSION=3.6
FROM golang:${GO_VERSION} AS dev-api

ENV GO111MODULE=on

WORKDIR /go/src/deviceAPI

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/deviceAPI 

EXPOSE 8086

FROM alpine:${ALPINE_VERSION} AS release-api
COPY --from=dev-api /go/bin/deviceAPI /smartac/deviceAPI
ENV PATH="/smartac:${PATH}"

ENTRYPOINT [ "deviceAPI" ]
