all: proxy

proxy:
	@echo "build tekton proxy"
	go build -o bin/tekton-proxy cmd/main.go

image:
	@echo "release tekton-proxy image"
	./build/build-image.sh