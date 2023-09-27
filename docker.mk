
DOCKER				?= docker
DOCKER_IMAGE 		?= pack-calc

docker-build:
	@$(DOCKER) build -t $(DOCKER_IMAGE) -f Dockerfile .

docker-run:
	@$(DOCKER) run -p 4000:4000 -p 4001:4001 -e ENVIRONMENT=develop $(DOCKER_IMAGE)

docker-build-and-run: docker-build docker-run