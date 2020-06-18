/* */

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var version = "v0.0.3"
var releaseDate = ""

type World struct {
	Conf     Configuration
	Camera   *Camera
	Entities map[string]*Entity
	Map      *GameMap
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
	}

	var player = &Entity{
		X:     w.Conf.MapWidth/2 - 1,
		Y:     w.Conf.MapHeight/2 - 1,
		Layer: 0,
		Char:  "@",
		Color: "cyan",
	}
	var npc = &Entity{
		X:     w.Conf.ScreenWidth/2 - 5,
		Y:     w.Conf.ScreenHeight/2 - 5,
		Layer: 0,
		Char:  "0",
		Color: "red",
	}
	w.Entities["player"] = player
	w.Entities["npc"] = npc

	w.Map.InitializeMap()
	//prettyPrintStruct(w)
	gameLoop(w)
}

func gameLoop(w *World) {
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
		renderAll(w)
	}

	fmt.Println("Closing Window and exit")
	closeWindow()
}

func handlePlayerAction(playerAction string, w *World) {
	player := w.Entities["player"]
	var dx, dy = 0, 0
	switch playerAction {
	case "up":
		dy = -1
	case "down":
		dy = 1
	case "left":
		dx = -1
	case "right":
		dx = 1
	case "upright":
		dx = 1
		dy = -1
	case "upleft":
		dx = -1
		dy = -1
	case "downright":
		dx = 1
		dy = 1
	case "downleft":
		dx = -1
		dy = 1
	}
	if !w.Map.Tiles[player.X+dx][player.Y+dy].Blocked {
		player.move(dx, dy)
	}
}
