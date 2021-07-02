BASE_IMAGE?=public.ecr.aws/micahhausler/alpine:3.14.0
IMAGE_REPO?=public.ecr.aws/micahhausler/samba
IMAGE_TAG?=4.14.15
IMAGE?=${IMAGE_REPO}:${IMAGE_TAG}

IMAGE_OUTPUT=--load
IMAGE_PLATFORM= linux/amd64
ifeq ($(PUBLISH),true)
	IMAGE_OUTPUT= --push
	IMAGE_PLATFORM= linux/aarch64,linux/amd64,linux/arm/v7
endif
.PHONY: image
image:
	docker buildx build \
		--platform=$(IMAGE_PLATFORM) \
		--build-arg BASE_IMAGE=${BASE_IMAGE} \
		--build-arg SAMBA_VERSION=${IMAGE_TAG} \
		$(IMAGE_OUTPUT) \
		--tag ${IMAGE} \
		.

.PHONY: all
all: image
