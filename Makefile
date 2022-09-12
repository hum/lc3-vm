.EXPORT_ALL_VARIABLES:
GOOS ?= $(uname -s)
GOARCH ?= $(uname -m)
LD_FLAGS := -ldflags="-s -w -X 'main.BuildDate=$(shell date)'"

run: build
	chmod +x bin/lc3_vm
	./bin/lc3_vm

build:
	go build -v -o bin/lc3_vm *.go

build_exec:
	env CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LD_FLAGS) -o bin/lc3_vm_opt *.go

