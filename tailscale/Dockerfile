ARG BASE_IMAGE=public.ecr.aws/micahhausler/alpine:3.13.5
FROM $BASE_IMAGE

MAINTAINER Micah Hausler <hausler.m@gmail.com>

ARG TAILSCALE_VERSION=1.6.0
ARG TARGETARCH

# Install samba
RUN apk --no-cache --no-progress upgrade && \
    apk --no-cache --no-progress add \
        iproute2 \
        tar \
        gzip \
        wget \
        tini \
    && rm -rf /tmp/* && \
    wget --progress dot:mega https://pkgs.tailscale.com/stable/tailscale_${TAILSCALE_VERSION}_${TARGETARCH}.tgz && \
    tar xvf tailscale_${TAILSCALE_VERSION}_${TARGETARCH}.tgz && \
    mv tailscale_${TAILSCALE_VERSION}_${TARGETARCH}/tailscaled /usr/sbin/tailscaled && \
    mv tailscale_${TAILSCALE_VERSION}_${TARGETARCH}/tailscale /usr/sbin/tailscale && \
    rm -rf tailscale_${TAILSCALE_VERSION}_${TARGETARCH} tailscale_${TAILSCALE_VERSION}_${TARGETARCH}.tgz

ENTRYPOINT ["/usr/sbin/tailscaled"]
