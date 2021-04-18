# Two-Stage Build
# Stage 1 - Prepare Binary File in Full Environment
FROM golang:latest AS builder

COPY . /go/src/github/quattad/mini-shopee/products-service
WORKDIR /go/src/github/quattad/mini-shopee/products-service
RUN set -x

RUN go get github/quattad/mini-shopee/products-service/src/main

WORKDIR /go/src/github/quattad/mini-shopee/products-service/src/main
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /products-service .

# CMD ["/products-service"]
# EXPOSE 8000

# Stage 2 - Take Only Binary File from /products-service
FROM scratch

COPY --from=builder /products-service .

ENV PORT 8000

# DEVELOPMENT
ENV ENV_SOURCE docker
ENV DB_DRIVER mysql
ENV DB_USER productsadm
ENV DB_PASSWORD Passwordvpc1!
ENV DB_HOSTNAME docker.for.win.localhost
ENV DB_PORT 3306
ENV DB_NAME products_service

ENTRYPOINT ["/products-service"]