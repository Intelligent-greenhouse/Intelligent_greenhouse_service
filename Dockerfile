FROM golang:1.21 AS builder

COPY . /src
WORKDIR /src

RUN mkdir -p /src/bin/
RUN GOPROXY=https://goproxy.cn go build -ldflags "-X main.Version=`git describe --tags --always`" -o ./bin/app ./cmd

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

ENV APP_CONFIG_DIR "/data/conf"

CMD ["./app"]
