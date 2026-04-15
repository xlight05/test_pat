FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /app /app
EXPOSE 9090
ENTRYPOINT ["/app"]
