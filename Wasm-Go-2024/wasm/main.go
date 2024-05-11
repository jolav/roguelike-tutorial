//go:build js && wasm

/* */

package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("hiFromWASM", js.FuncOf(hiFromWASM))
	select {}
}

func hiFromWASM(this js.Value, args []js.Value) any {
	return "Hi from WASM"
}
