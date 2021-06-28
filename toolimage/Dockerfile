ARG BASE_IMAGE=public.ecr.aws/micahhausler/alpine:3.13.5
FROM $BASE_IMAGE

MAINTAINER Micah Hausler <hausler.m@gmail.com>

RUN apk add --no-cache \
        bash \
        curl \
        drill \
        file \
        git \
        htop \
        iperf3 \
        iperf3-doc \
        iproute2 \
        jq \
        less \
        lsof \
        man-db \
        man-pages \
        net-tools \
        nmap \
        openssh \
        vim \
        wget \
        which

CMD ["/bin/bash"]
