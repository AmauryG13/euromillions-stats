FROM golang:1.17-alpine AS builder

ARG MODULE

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
COPY . /app

WORKDIR /app

RUN go mod download

RUN sh ./scripts/build.sh ${MODULE}

# Run container
FROM alpine:latest

ARG MODULE

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/${MODULE}/ems-${MODULE} .

CMD ["./ems-${MODULE}", "server"]