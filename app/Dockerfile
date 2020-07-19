FROM golang:alpine3.12 AS builder
WORKDIR /work
COPY ./ /work
RUN apk update && apk upgrade && \
    apk add make gcc g++ && \
    make gomod && \
    make build

FROM golang:alpine3.11
RUN apk update && apk upgrade
WORKDIR /work
COPY --from=builder /work/main /work/main
COPY --from=builder /work/.env /work/.env
CMD ./main
