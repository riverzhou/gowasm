package main

import (
	"fmt"
	"syscall/js"
)

const (
	width  = 640
	height = 480
)

var count int32 = 0
var count2 int32 = 0
var count3 int32 = 0

func getInt() int32 {
	count += 2
	return count
}

func getInt2() int32 {
	count2 += 10
	return count2
}

func getInt3() int32 {
	count3 += 10
	return count3
}

//export draw
func draw() {
	count = 0
	count2 = 0
	count3 = 0

	var canvas js.Value = js.
		Global().
		Get("document").
		Call("getElementById", "canvas")

	var context js.Value = canvas.Call("getContext", "2d")

	canvas.Set("height", height)
	canvas.Set("width", width)
	context.Call("clearRect", 0, 0, width, height)

	for i := 0; i < 30; i++ {
		tt := fmt.Sprintf(`rgba(%d, %d, %d, 0.5)`, 50+getInt(), 50+getInt(), 50+getInt())
		context.Call("beginPath")
		context.Set("strokeStyle", tt)
		context.Call("moveTo", getInt2(), getInt3())
		context.Call("lineTo", getInt2(), getInt3())
		context.Call("stroke")
	}

	context.Set("font", "30px Arial")
	context.Set("strokeStyle", "blue")
	context.Call("strokeText", "hello wasm", 100, 300)
}

func main() {
	println("Hello, Go WebAssembly!")
	println("canvas wasm loaded")
	println("go main finished")
}
