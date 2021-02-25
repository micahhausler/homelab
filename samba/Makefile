BASE_IMAGE?=public.ecr.aws/micahhausler/alpine:3.13.0
IMAGE_REPO?=public.ecr.aws/micahhausler/samba
IMAGE_TAG?=4.13.3
IMAGE?=${IMAGE_REPO}:${IMAGE_TAG}

image:
	docker build \
		--build-arg BASE_IMAGE=${BASE_IMAGE} \
		--build-arg SAMBA_VERSION=${IMAGE_TAG} \
		-t ${IMAGE} \
		.

push: image
	docker push ${IMAGE}

all: image

.PHONY: image all push
