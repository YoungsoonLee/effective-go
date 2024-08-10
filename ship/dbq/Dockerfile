# START: build
FROM golang:1.20-bookworm AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o ./dbq
# END: build

# START: app
FROM debian:bookworm-slim 
COPY --from=build /build/dbq /usr/local/bin

# Don't run as root user
RUN groupadd -r app && useradd --no-log-init -r -g app app
USER app

CMD dbq
# END: app

# example
# docker build . -t myimage
# docker run -it -p 80:80 -p 443:443 -e HELLO="hello world!" myimage