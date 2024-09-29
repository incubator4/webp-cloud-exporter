FROM golang:alpine as build

ARG GO111MODULE=on

WORKDIR /app
ADD go.* .
RUN go mod download
ADD . .
RUN go build -o /metrics-exporter main.go


FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
COPY --from=build /metrics-exporter /metrics-exporter
ENTRYPOINT ["/metrics-exporter"]