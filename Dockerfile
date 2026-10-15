FROM --platform=$BUILDPLATFORM golang:1.26.1-alpine AS builder

WORKDIR /ts2tg

RUN apk update && apk add --no-cache ca-certificates 

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o ./ts2tg



FROM scratch

LABEL org.opencontainers.image.title="ts2tg" \
    org.opencontainers.image.description="ts2tg notifies you on Telegram when users join/leave your TeamSpeak server" \
    org.opencontainers.image.authors="watn3y <me@watn3y.de>" \
    org.opencontainers.image.vendor="watn3y" \
    org.opencontainers.image.documentation="https://git.watn3y.de/watn3y/ts2tg/src/branch/master/README.md" \
    org.opencontainers.image.url="https://git.watn3y.de/watn3y/ts2tg" \
    org.opencontainers.image.source="https://git.watn3y.de/watn3y/ts2tg" \
    org.opencontainers.image.licenses="GPL-3.0-or-later"

WORKDIR /app

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /usr/share/ca-certificates /usr/share/ca-certificates

COPY --from=builder /ts2tg/ts2tg /app/ts2tg

ENTRYPOINT ["/app/ts2tg"]
