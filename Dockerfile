FROM golang:1.22-alpine as builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./exoplanets

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/exoplanets ./exoplanets

EXPOSE 6379

ENTRYPOINT ["./exoplanets", "-dbconn=my-redis:6379"]