FROM golang:1.21.1-bullseye AS builder

COPY . /dineflow-api-gateway
WORKDIR /dineflow-api-gateway
RUN go mod tidy
RUN go build

FROM debian:bullseye-slim
ENV GIN_MODE release

RUN mkdir /app
WORKDIR /app
COPY --from=builder /dineflow-api-gateway/dineflow-api-gateway /app

EXPOSE 8000

CMD ["/app/dineflow-api-gateway"]