/* */

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var version = "v0.0.6"
var releaseDate = ""

type World struct {
	Conf     Configuration
	Camera   *Camera
	Entities map[string]*Entity
	Map      *GameMap
	Fov      *FieldOfVision
}

func main() {
	checkFlags()
	runtime.LockOSThread()
	rand.Seed(time.Now().Unix())

	var w = new(World)

	loadConfigJSON(&w.Conf)
	initBLT(w.Conf)

	w = &World{
		Conf: w.Conf,
		Camera: &Camera{
			X:      0,
			Y:      0,
			Width:  w.Conf.CameraWidth,
			Height: w.Conf.CameraHeight,
		},
		Entities: make(map[string]*Entity),
		Map: &GameMap{
			Width:  w.Conf.MapWidth,
			Height: w.Conf.MapHeight,
		},
		Fov: &FieldOfVision{},
	}

	var player = &Entity{
		X:      w.Conf.MapWidth/2 - 1,
		Y:      w.Conf.MapHeight/2 - 1,
		Layer:  0,
		Char:   "@",
		Color:  "cyan",
		Facing: 'N',
	}
	/*var npc = &Entity{
		X:     w.Conf.ScreenWidth/2 - 5,
		Y:     w.Conf.ScreenHeight/2 - 5,
		Layer: 0,
		Char:  "0",
		Color: "red",
	}*/
	w.Entities["player"] = player
	//w.Entities["npc"] = npc

	w.Fov.initFOV()
	w.Map.initializeMap()
	//prettyPrintStruct(w)
	gameLoop(w)
}

func gameLoop(w *World) {
	w.Fov.rayCast(w)
	renderAll(w)

	for {
		key := readKeyboard()
		playerAction := keyToAction(key)
		//fmt.Println("Pressed ...", key, " - Action ...", playerAction)
		if playerAction == "exit" {
			break
		}
		if playerAction != "" {
			handlePlayerAction(playerAction, w)
		}
		w.Fov.rayCast(w)
		renderAll(w)
	}

	fmt.Println("Closing Window and exit")
	closeWindow()
}
