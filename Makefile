IMAGE_TAG=v1alpha1
QUAY_PASS?=biggestsecret
LIMITER_REQUESTS?=50
LIMITER_LIMIT?=300
LIMITER_BURST?=3
LIMITER_BURSTREPEAT?=3

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o burstlimiter .

dev: compile
	docker build -t quay.io/tamarakaufler/limiter:$(IMAGE_TAG) .

build: dev
	docker login quay.io -u tamarakaufler -p $(QUAY_PASS)
	docker push quay.io/tamarakaufler/limiter:$(IMAGE_TAG)

run:
	docker run \
	--name=limiter \
	--rm \
	quay.io/tamarakaufler/limiter:$(IMAGE_TAG) \
	-requests=$(LIMITER_REQUESTS) \
	-limit=$(LIMITER_LIMIT) \
	-burst=$(LIMITER_BURST) \
	-burstrepeat=$(LIMITER_BURSTREPEAT)
