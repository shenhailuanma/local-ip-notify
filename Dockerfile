ARG TARGETARCH
FROM alpine:3.20.3 AS base

WORKDIR /


# amd64 架构
FROM base AS amd64
COPY .build/linux/amd64/ip-notify /ip-notify


# arm64 架构
FROM base AS arm64
COPY .build/linux/arm64/ip-notify /ip-notify


# 选择最终阶段
FROM ${TARGETARCH} AS final
WORKDIR /

ENTRYPOINT ["/ip-notify"]
