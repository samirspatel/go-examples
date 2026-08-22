# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -trimpath -ldflags="-s -w" -o /hello

FROM scratch
COPY --from=build /hello /hello
ENTRYPOINT ["/hello"]
