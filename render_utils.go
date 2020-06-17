/* */

package main

import (
	blt "roguelike/bearlibterminal"
	"strconv"
)

func initBLT(c Configuration) {
	blt.Open()
	size := "size=" + strconv.Itoa(c.ScreenWidth) + "x"
	size += strconv.Itoa(c.ScreenHeight)
	title := "title='" + c.Title + "'"
	window := "window: " + size + "," + title
	fontSize := "size=" + strconv.Itoa(c.FontSize1)
	font := "font: " + c.Font1 + ", " + fontSize
	blt.Set(window + "; " + font)

	/*
		grey or gray, red, flame, orange, amber, yellow, lime, chartreuse, green, sea, turquoise, cyan, sky, azure, blue, han, violet, purple, fuchsia, magenta, pink, crimson, transparent
	*/
	blt.Set("palette.darkGrey=#636466")
	blt.Set("palette.choco=#bdb897")
	blt.Set("palette.purple=#979cbd")
	blt.Set("palette.violet=#b897db")
	blt.Set("palette.khaki=#d1cdb6")

	blt.BkColor(blt.ColorFromName("khaki"))
	blt.Clear()
}

func readKeyboard() (key int) {
	return blt.Read()
}

func closeWindow() {
	blt.Close()
}
