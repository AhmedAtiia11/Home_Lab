FROM golang:1.25-alpine AS build
WORKDIR /app
COPY main.go .
RUN go build -o app main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app .
EXPOSE 5000
CMD ["./app"]