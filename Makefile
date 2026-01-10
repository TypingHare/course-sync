EXECUTABLE := csync

.PHONY: build

build:
	mkdir -p bin
	go build -o bin/$(EXECUTABLE) ./cmd/csync/

clean:
	rm -rf bin
