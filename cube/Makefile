
all: cube 

cube:
	#tinygo build -no-debug -o cube.wasm -target cube.go
	#GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o cube.wasm cube.go
	GOOS=js GOARCH=wasm go build -o cube.wasm cube.go

clean:
	rm -rf cube.wasm

