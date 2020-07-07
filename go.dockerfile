# BUILD STAGE
FROM golang:alpine AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o goapp .

# RUN STAGE
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/templates /build/templates
COPY --from=builder /build/public /build/public
COPY --from=builder /build/goapp /build/goapp

WORKDIR /build

EXPOSE 8000
ENTRYPOINT ["/build/goapp"]