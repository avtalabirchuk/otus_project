FROM golang:1.16-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go build -o /out/app cmd/app/main.go

FROM alpine:latest
COPY --from=build /out/app .
COPY --from=build /src/configs/config.yml configs/
COPY --from=build /src/cache cache
CMD ["./app"]