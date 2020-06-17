/*
// LINUX
go build -ldflags="-X 'main.releaseDate=$(date -u +%F_%T)'"
// WIN 64
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc  go build -ldflags="-X 'main.when=$(date -u +%F_%T)'" -o roguelike.exe
*/

package main

type Configuration struct {
	ScreenWidth  int    `json:"screenWidth"`
	ScreenHeight int    `json:"screenHeight"`
	Title        string `json:"windowTitle"`
	Font1        string `json:"font1"`
	Font2        string `json:"font2"`
	FontSize1    int    `json:"fontSize1"`
	FontSize2    int    `json:"fontSize2"`
	MapWidth     int    `json:"mapWidth"`
	MapHeight    int    `json:"mapHeight"`
}

func getGlobalConfigJSON() (configjson []byte) {
	configjson = []byte(`
	{
		"screenWidth": 40,
		"screenHeight": 30,
		"windowTitle": "Roguelike Tutorial 2020",
		"font1": "assets/fonts/square.ttf",
		"font2": "assets/fonts/ubuntuMono.ttf",
		"fontSize1": 12,
		"fontSize2": 12,
		"mapWidth": 40,
		"mapHeight": 20
	}
	`)
	return
}
