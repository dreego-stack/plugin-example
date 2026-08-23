.PHONY: test test-race vet run clean

test:
	go test ./...

test-race:
	go test -race ./...

vet:
	go vet ./...

run:
	go run ./example

clean:
	rm -f *_dreego.go
	rm -rf bin/