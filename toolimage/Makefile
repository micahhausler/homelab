BASE_IMAGE?=public.ecr.aws/micahhausler/alpine:3.13.5
IMAGE_REPO?=public.ecr.aws/micahhausler/toolimage
IMAGE_TAG?=0.1
IMAGE?=${IMAGE_REPO}:${IMAGE_TAG}

image:
	docker build \
		--build-arg BASE_IMAGE=${BASE_IMAGE} \
		-t ${IMAGE} \
		.

push: image
	docker push ${IMAGE}

buildkit:
	buildctl \
		build \
		--frontend dockerfile.v0 \
		--opt platform=linux/aarch64,linux/amd64,linux/arm/v7 \
		--opt build-arg:BASE_IMAGE=$(BASE_IMAGE) \
		--local dockerfile=./ \
		--local context=. \
		--output type=image,oci-mediatypes=true,name=${IMAGE},push=true
		#--output type=oci,dest=/dev/null

all: image

.PHONY: image all push
