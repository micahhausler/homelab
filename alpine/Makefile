ALPINE_VERSION?=3.14
ALPINE_POINT?=0
PLATFORMS=x86_64 aarch64 armv7
ALPINE_TAR=alpine-minirootfs-${ALPINE_VERSION}.${ALPINE_POINT}-${PLATFORM}.tar.gz
ALPINE_SHA=${ALPINE_TAR}.sha256
URL=https://dl-cdn.alpinelinux.org/alpine/v${ALPINE_VERSION}/releases/${PLATFORM}/${ALPINE_TAR}
SHA_URL=${URL}.sha256

# Set to "true" to build/push multi-arch
PUBLISH?=

IMAGE_REPO?=public.ecr.aws/micahhausler/alpine
IMAGE?=${IMAGE_REPO}:${ALPINE_VERSION}.${ALPINE_POINT}

UNAME_S := $(shell uname -s)
SHASUM=sha256sum
ifeq ($(UNAME_S),Darwin)
	SHASUM = shasum -a 256
endif

.PHONY: get-shas
get-shas:
	wget \
		 --continue \
		--progress dot:mega \
		--max-redirect=0 \
		--domains dl-cdn.alpinelinux.org \
			$(foreach PLATFORM,$(PLATFORMS),$(SHA_URL))

.PHONY: get-tars
get-tars:
	wget \
		 --continue \
		--progress dot:mega \
		--max-redirect=0 \
		--domains dl-cdn.alpinelinux.org \
		$(foreach PLATFORM,$(PLATFORMS),$(URL))

.PHONY: verify-shas
verify-shas: get-tars get-shas
	$(SHASUM) -c *.sha256

.PHONY: prepare-tars
prepare-tars: verify-shas
	cp alpine-minirootfs-${ALPINE_VERSION}.${ALPINE_POINT}-x86_64.tar.gz alpine-minirootfs-${ALPINE_VERSION}.${ALPINE_POINT}-amd64.tar.gz
	cp alpine-minirootfs-${ALPINE_VERSION}.${ALPINE_POINT}-aarch64.tar.gz alpine-minirootfs-${ALPINE_VERSION}.${ALPINE_POINT}-arm64.tar.gz

IMAGE_OUTPUT=--load
IMAGE_PLATFORM= linux/amd64
ifeq ($(PUBLISH),true)
	IMAGE_OUTPUT= --push
	IMAGE_PLATFORM= linux/aarch64,linux/amd64,linux/arm/v7
endif
.PHONY: image
image: prepare-tars
	docker buildx build \
		--platform=$(IMAGE_PLATFORM) \
		--build-arg ALPINE_VERSION=$(ALPINE_VERSION) \
		--build-arg ALPINE_POINT=$(ALPINE_POINT) \
		$(IMAGE_OUTPUT) \
		--tag ${IMAGE} \
		.

.PHONY: clean
clean:
	rm alpine-minirootfs-*.tar.gz

.PHONY: all
all: image
