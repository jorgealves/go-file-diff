
.PHONY: build simpletest gotest

build:
	docker build -f build/docker/Dockerfile -t jorgeandrealves/go-file-diff .

simpletest: build
	docker run --rm -it jorgeandrealves/go-file-diff bash -c "go-file-diff ./examples/first.json ./examples/second.json"

gotest: build
	docker run --rm -it jorgeandrealves/go-file-diff bash -c "go test ./test"
