//go:build js && wasm

/* */

package main

type renderData struct {
	Nick string `json:"nick"`
}

func prepareRenderData(r run) renderData {
	rd := renderData{
		Nick: r.Nick,
	}
	return rd
}
