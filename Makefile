EXECUTABLE := csync

.PHONY: build test clean

build:
	@mkdir -p bin
	go build -o bin/$(EXECUTABLE) ./cmd/$(EXECUTABLE)/

test:
	GOCACHE=/tmp/go-build GOMODCACHE=/tmp/go-mod go test ./tests/...

# Build for Windows amd64.
build-windows-amd64:
	@mkdir -p bin
	GOOS=windows GOARCH=amd64 \
		 go build -o bin/$(EXECUTABLE)_amd64.exe ./cmd/$(EXECUTABLE)/

# Build for Windows arm64.
build-windows-arm64:
	@mkdir -p bin
	GOOS=windows GOARCH=arm64 \
		 go build -o bin/$(EXECUTABLE)_arm64.exe ./cmd/$(EXECUTABLE)/

# Build for Linux amd64.
build-linux-amd64:
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 \
		 go build -o bin/$(EXECUTABLE)_linux_amd64 ./cmd/$(EXECUTABLE)/

# Build for Linux arm64.
build-linux-arm64:
	@mkdir -p bin
	GOOS=linux GOARCH=arm64 \
		 go build -o bin/$(EXECUTABLE)_linux_arm64 ./cmd/$(EXECUTABLE)/

# Build for Mac amd64.
build-mac-amd64:
	@mkdir -p bin
	GOOS=darwin GOARCH=amd64 \
		 go build -o bin/$(EXECUTABLE)_mac_amd64 ./cmd/$(EXECUTABLE)/

# Build for Mac arm64.
build-mac-arm64:
	@mkdir -p bin
	GOOS=darwin GOARCH=arm64 \
		 go build -o bin/$(EXECUTABLE)_mac_arm64 ./cmd/$(EXECUTABLE)/

# Build for all platforms
build-all:
	make build-windows-amd64
	make build-windows-arm64
	make build-linux-amd64
	make build-linux-arm64
	make build-mac-amd64
	make build-mac-arm64

clean:
	rm -rf bin
