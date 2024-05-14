//go:build js && wasm

/* */

package main

import (
	"strconv"
	"strings"
	"syscall/js"
)

type config struct {
	Version     string `json:"version"`
	NicksFile   string `json:"nicksFile"`
	TokenLen    int    `json:"tokenLength"`
	NicChars    int    `json:"nickChars"`
	NickNumbers int    `json:"nickNumbers"`
}

type camera struct {
	Cols int
	Rows int
}

type app struct {
	C   config `json:"config"`
	r   run
	cam camera
}

func main() {
	// Load Conf
	var a = new(app)
	a.loadConfigJSON()

	params := strings.Split(js.Global().Call("getViewSize").String(), "_")
	cols, _ := strconv.Atoi(params[0])
	rows, _ := strconv.Atoi(params[1])
	a.cam = camera{
		Cols: cols,
		Rows: rows,
	}
	a.r = *newRun(*a)

	js.Global().Set("getApiVersion", js.FuncOf(a.apiVersion))
	js.Global().Set("getNewTurn", js.FuncOf(a.apiNewTurn))

	//var obj = structToJSObject(a.r)
	//js.Global().Set("myRun", obj)
	//js.Global().Call("init2", obj)
	js.Global().Call("allInitialized")

	select {}
}
