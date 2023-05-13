IMAGE=github.com/tobbstr-examples/client-server-monorepo:v1.0.0

# Build proto-generation docker image.
proto_image:
	DOCKER_BUILDKIT=1 docker build -f build/Dockerfile -t $(IMAGE) .

# Run go generate in docker.
gen: proto_image
	 docker run --rm -w /wd/in -v $(CURDIR):/wd/in $(IMAGE) go generate ./...
# cannot generate files outside of 'server' Go module so they are generated inside it and then moved
# to the shared module.
	 @cp pb/v1/* ../shared/pb/v1
	 @rm -rf pb
