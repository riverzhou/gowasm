
all: test

test:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o test.wasm test.go

clean:
	rm -rf test.wasm

