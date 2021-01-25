ARG BASE_IMAGE=public.ecr.aws/s3t5k9h7/alpine:3.13.0
FROM $BASE_IMAGE

MAINTAINER Micah Hausler <hausler.m@gmail.com>

ARG TAILSCALE_VERSION=1.2.10
ARG ARCH=amd64

# Install samba
RUN apk --no-cache --no-progress upgrade && \
    apk --no-cache --no-progress add \
        iproute2 \
        tar \
        gzip \
        wget \
        tini \
    && rm -rf /tmp/* && \
    wget https://pkgs.tailscale.com/stable/tailscale_${TAILSCALE_VERSION}_${ARCH}.tgz && \
    tar xvf tailscale_${TAILSCALE_VERSION}_${ARCH}.tgz && \
    mv tailscale_${TAILSCALE_VERSION}_${ARCH}/tailscaled /usr/sbin/tailscaled && \
    mv tailscale_${TAILSCALE_VERSION}_${ARCH}/tailscale /usr/sbin/tailscale && \
    rm -rf tailscale_${TAILSCALE_VERSION}_${ARCH} tailscale_${TAILSCALE_VERSION}_${ARCH}.tgz

ENTRYPOINT ["/usr/sbin/tailscaled"]
