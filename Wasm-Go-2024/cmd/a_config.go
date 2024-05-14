//go:build js && wasm

/* */

package main

import (
	"encoding/json"
	"log"
)

func getConfigJSON() (configjson []byte) {

	configjson = []byte(`
	{
		"config": {
			"version": "0.0.1",
			"nicksFile": "./assets/surname.txt",
			"tokenLength": 40,
			"nickChars": 6,
			"nickNumbers": 2
		}
	}
	`)

	return
}

func (a *app) loadConfigJSON() {
	err := json.Unmarshal(getConfigJSON(), a)
	if err != nil {
		log.Fatal("Error parsing JSON config => \n", err)
	}
}
