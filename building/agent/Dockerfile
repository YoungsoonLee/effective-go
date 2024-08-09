FROM golang:1.20-bookworm AS build

ENV CGO_ENABLED=0
WORKDIR /app
COPY go.mod .
COPY agent.go .
RUN go build -o agent .

FROM alpine:3.18
COPY --from=build /app/agent /usr/local/bin
CMD ["/usr/local/agent"]
