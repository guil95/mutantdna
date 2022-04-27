FROM golang:buster as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE='on'

WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o /usr/local/bin/mutant ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /usr/local/bin/mutant .

EXPOSE 80
CMD ["./mutant"]