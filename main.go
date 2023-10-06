package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	// Expose the callJS function to JavaScript
	js.Global().Set("callGoFetch", js.FuncOf(jsCallJS))

	<-c
}

func jsCallJS(this js.Value, args []js.Value) interface{} {
	jsFunc := js.Global().Get("fetchFromGo")
	promise := jsFunc.Invoke(args[0].String())

	// Return the promise directly
	return promise
}
