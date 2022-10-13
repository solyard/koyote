FROM golang:alpine as build-env
LABEL maintainer="dizstorm@gmail.com"
COPY . /app
WORKDIR /app
RUN go mod download && CGO_ENABLED=0 go build -o /usr/bin/koyote .

FROM alpine as final
COPY --from=build-env /usr/bin/koyote /koyote
ENTRYPOINT ["/koyote"]