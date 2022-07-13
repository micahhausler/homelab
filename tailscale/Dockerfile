ARG BASE_IMAGE=public.ecr.aws/micahhausler/alpine:3.13.0
FROM $BASE_IMAGE

MAINTAINER Micah Hausler <hausler.m@gmail.com>

ARG TARGETARCH
ARG TAILSCALE_VERSION

# Install samba
RUN apk --no-cache --no-progress upgrade && \
    apk --no-cache --no-progress add \
        iproute2 \
        tar \
        gzip \
        wget \
        tini \
    && rm -rf /tmp/* && \
    wget https://pkgs.tailscale.com/stable/tailscale_${TAILSCALE_VERSION}_${TARGETARCH}.tgz && \
    tar xvf tailscale_${TAILSCALE_VERSION}_${TARGETARCH}.tgz && \
    mv tailscale_${TAILSCALE_VERSION}_${TARGETARCH}/tailscaled /usr/sbin/tailscaled && \
    mv tailscale_${TAILSCALE_VERSION}_${TARGETARCH}/tailscale /usr/sbin/tailscale && \
    rm -rf tailscale_${TAILSCALE_VERSION}_${TARGETARCH} tailscale_${TAILSCALE_VERSION}_${TARGETARCH}.tgz && \
    mkdir -p /var/lib/tailscale/

ENTRYPOINT ["/usr/sbin/tailscaled"]
