ARG BASE_IMAGE
FROM $BASE_IMAGE

MAINTAINER Micah Hausler <hausler.m@gmail.com>

ARG AVAHI_VERSION
ARG AVAHI_RELEASE=r2

RUN apk add --no-cache \
        avahi=$AVAHI_VERSION-$AVAHI_RELEASE \
        avahi-tools=$AVAHI_VERSION-$AVAHI_RELEASE \
        avahi-doc=$AVAHI_VERSION-$AVAHI_RELEASE \  
        less \
        man-db \
        man-pages && \
    rm /etc/avahi/services/sftp-ssh.service && \
    rm /etc/avahi/services/ssh.service

ADD avahi-daemon.conf /etc/avahi/avahi-daemon.conf
ADD smb.service /etc/avahi/services/smb.service

CMD ["/usr/sbin/avahi-daemon"]
