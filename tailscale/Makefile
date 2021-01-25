BASE_IMAGE?=public.ecr.aws/s3t5k9h7/alpine:3.13.0
IMAGE_REPO?=public.ecr.aws/s3t5k9h7/tailscale
IMAGE_TAG?=v1.2.10
IMAGE?=${IMAGE_REPO}:${IMAGE_TAG}

image:
	docker build \
		--build-arg BASE_IMAGE=${BASE_IMAGE} \
		-t ${IMAGE} \
		.

push: image
	docker push ${IMAGE}

all: image

.PHONY: image all push
