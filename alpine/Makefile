ALPINE_VERSION ?= 3.16
ALPINE_POINT   ?= 0

ARCHES = x86_64 aarch64

IMAGE_REPO ?= public.ecr.aws/micahhausler/alpine
IMAGE      ?= $(IMAGE_REPO):$(ALPINE_VERSION).$(ALPINE_POINT)

TARBALL = alpine-minirootfs-$(ALPINE_VERSION).$(ALPINE_POINT)-$(ARCH).tar.gz
DIGEST  = $(TARBALL).sha256
URL     = https://dl-cdn.alpinelinux.org/alpine/v$(ALPINE_VERSION)/releases/$(ARCH)/$(TARBALL)

define ARCH_RULES
ARCH = $(1)

# target-specific variable overrides
$$(subst x86_,amd,$$(subst aarch,arm,$$(TARBALL))) $$(TARBALL) $$(DIGEST): TARBALL=alpine-minirootfs-$(ALPINE_VERSION).$(ALPINE_POINT)-$(ARCH).tar.gz
$$(TARBALL) $$(DIGEST): DIGEST=$$(TARBALL).sha256
$$(TARBALL) $$(DIGEST): URL = https://dl-cdn.alpinelinux.org/alpine/v$(ALPINE_VERSION)/releases/$(ARCH)/$$(TARBALL)

$$(TARBALL): $$(DIGEST)
	wget \
		--progress dot:mega \
		--max-redirect=0 \
		--domains dl-cdn.alpinelinux.org \
		$$(URL)
	@touch $$(TARBALL)
	sha256sum -c $$(DIGEST)

# docker-friendly arch name
$$(subst x86_,amd,$$(subst aarch,arm,$$(TARBALL))): $$(TARBALL)
	if [ -f $$@ ]; then rm $$@; fi
	ln -s $$(TARBALL) $$@

$$(DIGEST): 
	wget \
		--progress dot \
		--max-redirect=0 \
		--domains dl-cdn.alpinelinux.org \
		$$(URL).sha256
	@touch $$(DIGEST)

endef
$(foreach ARCH, $(ARCHES), $(eval $(call ARCH_RULES,$(ARCH)) ))

TARBALLS    = $(foreach ARCH,$(ARCHES),$(TARBALL))
DIGESTS     = $(foreach TAR,$(TARBALLS),$(TAR).sha256)
DOCKER_TARS = $(foreach TARBALL,$(TARBALLS),$(subst x86_,amd,$(subst aarch,arm,$(TARBALL))))
$(DOCKER_TARS): $(DIGESTS)

tarballs: $(DOCKER_TARS)

define DOCKER_BUILD
# $1 = arches
# $2 = alpine version
# $3 = output format
docker buildx build \
	--platform $(1) \
	--build-arg ALPINE_VERSION=$(2) \
	--tag $(IMAGE) \
	-o type=$(3) \
	.
endef

.PHONY: image
image: DOCKER_ARCHES = $(subst x86_,amd,linux/$(subst aarch,arm,$(shell uname -m)))
image: OUTPUT=docker
image: $(DOCKER_TARS)
	$(call DOCKER_BUILD,$(DOCKER_ARCHES),$(ALPINE_VERSION).$(ALPINE_POINT),$(OUTPUT))

null  :=
space := $(null) #
comma := ,
.PHONY: push
push: DOCKER_ARCHES = $(subst $(space),$(comma),$(foreach ARCH,$(ARCHES),$(subst x86_,amd,linux/$(subst aarch,arm,$(ARCH)))))
push: OUTPUT=registry
push:
	$(call DOCKER_BUILD,$(DOCKER_ARCHES),$(ALPINE_VERSION).$(ALPINE_POINT),$(OUTPUT))

.PHONY: clean
clean:
	rm alpine-minirootfs-*.tar.gz

.PHONY: all
all: image
