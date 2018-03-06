all: test install
install:
	go install ./...
test:
	go test ./...
deps:
	hash dep 2>/dev/null || { echo >&2 "Golang dep tool has to be installed. Aborting."; exit 1; }
	dep ensure
