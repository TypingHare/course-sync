EXECUTABLE := csync

.PHONY: build

build:
	mkdir -p bin
	go build -o bin/$(EXECUTABLE) ./cmd/$(EXECUTABLE)/

clean:
	rm -rf bin
