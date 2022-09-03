REPOSITORY_NAME ?= technicaltest
IMAGE_NAME ?= usedeall
DEPLOY_DOCKER_IMAGE_URL ?= technicaltest/usedeall
.PHONY: docker-build-local docker-run-prod docker-run-dev

run:
	@go run main.go

docker-build-local-dev:
	docker build -t ${REPOSITORY_NAME}/${IMAGE_NAME}:dev --build-arg SSH_PRIVATE_KEY="$$(cat ~/.ssh/id_rsa)" .	

docker-build-local-prod:
	docker build -t ${REPOSITORY_NAME}/${IMAGE_NAME}:latest --build-arg SSH_PRIVATE_KEY="$$(cat ~/.ssh/id_rsa)" .		

docker-run-dev:
	docker run --rm -it -p 3000:3000 -v ${PWD}:/go/src/usedeall ${REPOSITORY_NAME}/${IMAGE_NAME}:dev	