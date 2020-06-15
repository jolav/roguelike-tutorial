/* */

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"

	blt "roguelike/bearlibterminal"
)

var version = "v0.0.1"
var releaseDate = ""

type Player struct {
	X int
	Y int
}

func main() {
	checkFlags()
	runtime.LockOSThread()
	rand.Seed(time.Now().Unix())

	var c Configuration
	loadConfigJSON(&c)
	//prettyPrintStruct(c)
	initBLT(c)
	var p = &Player{
		X: c.ScreenWidth/2 - 1,
		Y: c.ScreenHeight/2 - 1,
	}
	gameLoop(p, c)
}

func initBLT(c Configuration) {
	blt.Open()
	size := "size=" + strconv.Itoa(c.ScreenWidth) + "x"
	size += strconv.Itoa(c.ScreenHeight)
	title := "title='" + c.Title + "'"
	window := "window: " + size + "," + title
	fontSize := "size=" + strconv.Itoa(c.FontSize1)
	font := "font: " + c.Font1 + ", " + fontSize
	blt.Set(window + "; " + font)
	blt.Clear()
}

func gameLoop(p *Player, c Configuration) {
	blt.Clear()
	drawPlayer(p)

	for {
		blt.Refresh()
		key := blt.Read()
		playerAction := keyToAction(key)
		fmt.Println("Pressed ...", key, " - Action ...", playerAction)
		if playerAction == "exit" {
			break
		}
		if playerAction != "" {
			handlePlayerAction(playerAction, p, c)
		}
		blt.Clear()
		drawPlayer(p)
	}

	fmt.Println("Closing Window and exit")
	blt.Close()
}

func drawPlayer(p *Player) {
	blt.Layer(0)
	blt.Color(blt.ColorFromName("cyan"))
	blt.Print(p.X, p.Y, "@")

}

func handlePlayerAction(playerAction string, p *Player, c Configuration) {
	oldX := p.X
	oldY := p.Y
	switch playerAction {
	case "up":
		p.Y--
	case "down":
		p.Y++
	case "left":
		p.X--
	case "right":
		p.X++
	case "upright":
		p.X++
		p.Y--
	case "upleft":
		p.X--
		p.Y--
	case "downright":
		p.X++
		p.Y++
	case "downleft":
		p.X--
		p.Y++
	}
	if p.X > c.ScreenWidth-1 || p.X < 0 {
		p.X = oldX
	}
	if p.Y > c.ScreenHeight-1 || p.Y < 0 {
		p.Y = oldY
	}

}
