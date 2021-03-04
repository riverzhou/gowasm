
all: lute 

lute:
	#tinygo build -no-debug -o lutewasm.wasm -target wasm lute   
	GOOS=js GOARCH=wasm go build -o lutewasm.wasm lute  

clean:
	rm -rf lutewasm.wasm

