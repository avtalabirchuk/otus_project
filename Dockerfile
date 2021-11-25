FROM golang:1.16-alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN go mod download
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/app cmd/app/main.go

FROM alpine
COPY --from=build /out/app .
COPY --from=build /src/configs/config.yml configs/
COPY --from=build /src/cache cache
CMD ["./app", "--config", "configs/config.yml"]