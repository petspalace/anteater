all: local containers

local:
	GOOS="linux" GOARCH="amd64" go build -o bin/anteater-linux-amd64 ./cmd/anteater
	GOOS="linux" GOARCH="arm64" go build -o bin/anteater-linux-arm64 ./cmd/anteater
	GOOS="freebsd" GOARCH="amd64" go build -o bin/anteater-freebsd-amd64 ./cmd/anteater
	GOOS="freebsd" GOARCH="arm64" go build -o bin/anteater-freebsd-arm64 ./cmd/anteater
containers:
	podman build --jobs=2 --platform=linux/amd64,linux/arm64 --manifest anteater ./cmd/anteater
