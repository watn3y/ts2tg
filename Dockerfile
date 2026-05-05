FROM --platform=$BUILDPLATFORM golang:1.26.1-alpine AS builder

WORKDIR /ts2tg

RUN apk update && apk add --no-cache ca-certificates 

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o ./ts2tg



FROM scratch
LABEL org.opencontainers.image.source=https://git.watn3y.de/watn3y/ts2tg
LABEL org.opencontainers.image.licenses=GPL-3.0
WORKDIR /app

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /usr/share/ca-certificates /usr/share/ca-certificates

COPY --from=builder /ts2tg/ts2tg /app/ts2tg

ENTRYPOINT ["/app/ts2tg"]
