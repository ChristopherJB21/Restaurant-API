FROM golang:alpine AS build-stage

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /d3golang

FROM gcr.io/distroless/static-debian12:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /d3golang /d3golang
COPY --from=build-stage /app/api.config api.config

ENTRYPOINT [ "/d3golang" ]