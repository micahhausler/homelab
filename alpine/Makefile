ALPINE_VERSION=3.13
ALPINE_POINT=0
PLATFORM=x86_64
ALPINE_TAR=alpine-minirootfs-${ALPINE_VERSION}.${ALPINE_POINT}-${PLATFORM}.tar.gz
URL=https://dl-cdn.alpinelinux.org/alpine/v${ALPINE_VERSION}/releases/${PLATFORM}/${ALPINE_TAR}
SHA_URL=${URL}.sha256

IMAGE_REPO?=public.ecr.aws/micahhausler/alpine
IMAGE?=${IMAGE_REPO}:${ALPINE_VERSION}.${ALPINE_POINT} 

${ALPINE_TAR}.sha256:
	wget \
		--progress dot \
		--max-redirect=0 \
		--domains dl-cdn.alpinelinux.org \
		${SHA_URL}

verify-sha: ${ALPINE_TAR}.sha256
	sha256sum -c ${ALPINE_TAR}.sha256

${ALPINE_TAR}:
	wget \
		--progress dot:mega \
		--max-redirect=0 \
		--domains dl-cdn.alpinelinux.org \
		${URL}

image: ${ALPINE_TAR} verify-sha
	docker build \
		--build-arg TARBALL=${ALPINE_TAR} \
		-t ${IMAGE} \
		.

clean:
	rm  ${ALPINE_TAR}

push: image
	docker push ${IMAGE}

all: image

.PHONY: image all verify-sha clean push
