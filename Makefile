
.PHONY: build
build:
	@echo "+ $@"
	CGO_ENABLED=0 go build -o bin/server \
        -ldflags "-w -s" .
