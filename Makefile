LDFLAGS                 := -ldflags "-w -s --extldflags '-static'"



build:
	$(info Building auth-service)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/auth-service ${LDFLAGS} ./cmd/main.go

test:
	$(info Testing auth-service)
	go test -v ./cmd

start:
	$(info Run!)
	docker-compose -f docker-compose.yaml up -d

stop:
	$(info stopping VC)
	docker-compose -f docker-compose.yaml rm -s -f

hard_restart: stop start

restart:
	docker restart auth-service

get_release-tag:
	@date +'%Y%m%d%H%M%S%9N'

ifndef VERSION
VERSION := latest
endif

DOCKER_TAG_AUTH_SERVICE := docker.sunet.se/auth-service/auth-service:$(VERSION)

docker-build:
	$(info Docker Building with tag: $(VERSION))
	docker build --tag $(DOCKER_TAG_AUTH_SERVICE) --file dockerfiles/issuer .


docker-push:
	$(info Pushing docker images)
	docker push $(DOCKER_TAG_AUTH_SERVICE)

ci_build: docker-build_auth-service docker-push
	$(info CI Build)

release-tag:
	git tag -s ${RELEASE} -m"release ${RELEASE}"

release_push:
	git push --tags

release: release-tag release_push
	$(info making release ${RELEASE})