FROM golang:alpine AS build-stage

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /restaurant

FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /restaurant /restaurant
COPY --from=build-stage /app/api.config api.config
COPY --from=build-stage /app/publicKey publicKey

ENTRYPOINT [ "/restaurant" ]