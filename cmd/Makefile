all: sandbox

sandbox:
	@echo "build sandbox"
	go build -tags example -o bin/sandbox main.go

image:
	@echo "release sandbox image"
	./build/build-image.sh
