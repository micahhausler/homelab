BASE_IMAGE ?= public.ecr.aws/micahhausler/alpine:3.16.0
IMAGE_REPO ?= public.ecr.aws/micahhausler/tailscale
IMAGE_TAG  ?= v1.26.2
IMAGE      ?= ${IMAGE_REPO}:${IMAGE_TAG}

ARCHES = amd64 arm64

define DOCKER_BUILD
# $1 = arches
# $2 = tailscale version
# $3 = output format
docker buildx build \
	--platform $(1) \
	--build-arg TAILSCALE_VERSION=$(2) \
	--build-arg BASE_IMAGE=$(BASE_IMAGE) \
	--tag $(IMAGE) \
	-o type=$(3) \
	.
endef

null  :=
space := $(null) #
comma := ,

.PHONY: image 
image: DOCKER_ARCHES = $(addprefix linux/,$(shell uname -m))
image: OUTPUT        = docker
image: V             = $(subst v,,$(IMAGE_TAG))
#push image: 
	#$(call DOCKER_BUILD,$(DOCKER_ARCHES),$(V),$(OUTPUT))

.PHONY: push
push: DOCKER_ARCHES = $(subst $(space),$(comma),$(addprefix linux/,$(ARCHES)))
push: OUTPUT        = registry
push: V             = $(subst v,,$(IMAGE_TAG))

image push: 
	$(call DOCKER_BUILD,$(DOCKER_ARCHES),$(V),$(OUTPUT))

.PHONY: all
all: image

