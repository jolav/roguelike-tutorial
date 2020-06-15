/* */

package main

import blt "roguelike/bearlibterminal"

func keyToAction(key int) (action string) {
	switch key {
	// exit
	case blt.TK_CLOSE, blt.TK_ESCAPE:
		return "exit"
	// movement
	case blt.TK_RIGHT, blt.TK_KP_6:
		return "right"
	case blt.TK_LEFT, blt.TK_KP_4:
		return "left"
	case blt.TK_UP, blt.TK_KP_8:
		return "up"
	case blt.TK_DOWN, blt.TK_KP_2:
		return "down"
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
