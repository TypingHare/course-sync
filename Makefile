.PHONY: build

build:
	go build -o build/course-sync .
	mkdir -p bin
	ln -sf $(PWD)/build/course-sync bin/csync
