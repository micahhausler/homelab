
ifdef MAKECMDGOALS
TARGET=$(MAKECMDGOALS)
else
TARGET=$(DEFAULT_GOAL)
endif

all: makes

image: makes

push: login makes

makes:
	make -C alpine $(TARGET)
	make -C avahi $(TARGET)
	make -C samba $(TARGET)

login:
	aws ecr-public get-login-password \
		--region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws

.PHONY: all makes image login push
