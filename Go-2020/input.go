/* */

package main

import (
	blt "roguelike/bearlibterminal"
)

func handlePlayerAction(playerAction string, w *World) {
	player := w.Entities["player"]
	var dx, dy = 0, 0
	//fmt.Println(string(player.Facing))
	switch playerAction {
	case "advance":
		switch player.Facing {
		case 'N':
			dy = -1
		case 'E':
			dx = 1
		case 'S':
			dy = 1
		case 'W':
			dx = -1
		}
	case "strafeL":
		switch player.Facing {
		case 'N':
			dx = -1
		case 'E':
			dy = -1
		case 'S':
			dx = 1
		case 'W':
			dy = 1
		}
	case "strafeR":
		switch player.Facing {
		case 'N':
			dx = 1
		case 'E':
			dy = 1
		case 'S':
			dx = -1
		case 'W':
			dy = -1
		}
	case "backpedal":
		switch player.Facing {
		case 'N':
			dy = 1
		case 'E':
			dx = -1
		case 'S':
			dy = -1
		case 'W':
			dx = 1
		}
	case "turnL":
		switch player.Facing {
		case 'N':
			player.Facing = 'W'
		case 'E':
			player.Facing = 'N'
		case 'S':
			player.Facing = 'E'
		case 'W':
			player.Facing = 'S'
		}
	case "turnR":
		switch player.Facing {
		case 'N':
			player.Facing = 'E'
		case 'E':
			player.Facing = 'S'
		case 'S':
			player.Facing = 'W'
		case 'W':
			player.Facing = 'N'
		}
	}
	if !w.Map.Tiles[player.X+dx][player.Y+dy].Blocked {
		player.move(dx, dy)
	}
}

func keyToAction(key int) (action string) {
	switch key {
	// exit
	case blt.TK_CLOSE, blt.TK_ESCAPE:
		return "exit"
	// movement
	case blt.TK_RIGHT, blt.TK_KP_6:
		if blt.Check(blt.TK_ALT) == 1 {
			return "strafeR"
		}
		return "turnR" //"right"
	case blt.TK_LEFT, blt.TK_KP_4:
		if blt.Check(blt.TK_ALT) == 1 {
			return "strafeL"
		}
		return "turnL" //"left"
	case blt.TK_UP, blt.TK_KP_8:
		return "advance" //"up"
	case blt.TK_DOWN, blt.TK_KP_2:
		return "backpedal" //"down"
	case blt.TK_KP_9:
		return "upright"
	case blt.TK_KP_7:
		return "upleft"
	case blt.TK_KP_3:
		return "downright"
	case blt.TK_KP_1:
		return "downleft"
	}
	return ""
}
