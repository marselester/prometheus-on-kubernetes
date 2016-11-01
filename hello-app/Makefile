# Default hello app version.
v ?= v3

# Build a static binary so the app is self-contained.
build:
	GOOS=linux go build \
		--tags netgo \
		--ldflags '-extldflags "-static"' \
		-o hello ./$(v)/

docker_build:
	sudo docker build -t marselester/prom-on-k8s:$(v) .

docker_push:
	sudo docker push marselester/prom-on-k8s:$(v)
