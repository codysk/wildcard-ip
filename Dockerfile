FROM golang:1.17-alpine AS build-env
RUN apk add --update make git gcc musl-dev
ADD . /project
WORKDIR /project
RUN make

FROM alpine:3.9

COPY --from=build-env /project/bin/wildcard-ip /app/wildcard-ip
WORKDIR /app
ENTRYPOINT ["./wildcard-ip"]
