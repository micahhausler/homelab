ARG BASE_IMAGE=public.ecr.aws/s3t5k9h7/alpine:3.13.0
FROM $BASE_IMAGE

MAINTAINER Micah Hausler <hausler.m@gmail.com>

ARG SAMBA_VERSION=4.13.3
ARG SAMBA_RELEASE=r1

# Install samba
RUN apk --no-cache --no-progress upgrade && \
    apk --no-cache --no-progress add \
        bash \
        less \
        man-db \
        man-pages \
        samba=$SAMBA_VERSION-$SAMBA_RELEASE \
        samba-doc=$SAMBA_VERSION-$SAMBA_RELEASE \
        shadow \
        tini \
        tzdata && \
    addgroup -S smb && \
    adduser -S -D -H -h /tmp -s /sbin/nologin -G smb -g 'Samba User' smbuser && \
    rm -rf /tmp/*

COPY smb.conf /etc/samba/smb.conf

EXPOSE 137/udp 138/udp 139 445

HEALTHCHECK --interval=60s --timeout=15s \
            CMD smbclient -L \\localhost -U % -m SMB3

VOLUME [ \
    "/etc", \
    "/var/cache/samba", \
    "/var/lib/samba", \
    "/var/log/samba", \
    "/run/samba"]

# TODO: set USER 
ENTRYPOINT ["/sbin/tini", "--", "smbd", "-FS", "--no-process-group"]
