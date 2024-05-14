//go:build js && wasm

/* */

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func (a app) apiVersion(this js.Value, args []js.Value) any {
	return a.C.Version
}

func (a app) apiNewTurn(this js.Value, args []js.Value) any {
	rd := prepareRenderData(a.r)
	return structToJSObject(rd)
}

func structToJSObject(v any) any {
	aux, err := json.Marshal(v)
	if err != nil {
		fmt.Println("Error StructToObjectJS => ", err)
		return err
	}
	obj := js.Global().Get("JSON").Call("parse", string(aux))
	return obj
}
