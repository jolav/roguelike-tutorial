/* */

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var version = "v0.0.2"
var releaseDate = ""

type World struct {
	Conf     Configuration
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
		Conf:     w.Conf,
		Entities: make(map[string]*Entity),
		Map: &GameMap{
			Width:  w.Conf.MapWidth,
			Height: w.Conf.MapHeight,
		},
	}

	var player = &Entity{
		X:     w.Conf.ScreenWidth/2 - 1,
		Y:     w.Conf.ScreenHeight/2 - 1,
		Char:  "@",
		Color: "cyan",
	}
	var npc = &Entity{
		X:     w.Conf.ScreenWidth/2 - 5,
		Y:     w.Conf.ScreenHeight/2 - 5,
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
		fmt.Println("Pressed ...", key, " - Action ...", playerAction)
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
	oldX := player.X
	oldY := player.Y
	switch playerAction {
	case "up":
		player.move(0, -1)
	case "down":
		player.move(0, 1)
	case "left":
		player.move(-1, 0)
	case "right":
		player.move(1, 0)
	case "upright":
		player.move(1, -1)
	case "upleft":
		player.move(-1, -1)
	case "downright":
		player.move(1, 1)
	case "downleft":
		player.move(-1, 1)
	}
	if w.Map.Tiles[player.X][player.Y].Blocked {
		player.X = oldX
		player.Y = oldY
	}
}
