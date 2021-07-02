BASE_IMAGE?=public.ecr.aws/micahhausler/alpine:3.14.0
IMAGE_REPO?=public.ecr.aws/micahhausler/tailscale
TAILSCALE_VERSION?=1.10.0
IMAGE_TAG?=v${TAILSCALE_VERSION}
IMAGE?=${IMAGE_REPO}:${IMAGE_TAG}

IMAGE_OUTPUT=--load
IMAGE_PLATFORM=linux/amd64
ifeq ($(PUBLISH),true)
	IMAGE_OUTPUT=--push
	IMAGE_PLATFORM=linux/aarch64,linux/amd64,linux/arm/v7
endif
.PHONY: image
image:
	docker buildx build \
		--platform=$(IMAGE_PLATFORM) \
		--build-arg BASE_IMAGE=$(BASE_IMAGE) \
		--build-arg TAILSCALE_VERSION=$(TAILSCALE_VERSION) \
		$(IMAGE_OUTPUT) \
		--tag ${IMAGE} \
		.

.PHONY: all
all: image
