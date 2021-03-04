
all: mouse

mouse:
	#tinygo build -no-debug -o mouse.wasm -target wasm mouse.go
	GOOS=js GOARCH=wasm go build -o mouse.wasm  mouse  

clean:
	rm -rf mouse.wasm

