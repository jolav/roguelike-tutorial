![Version](https://img.shields.io/badge/version-0.0.3-orange.svg)  
![Maintained YES](https://img.shields.io/badge/Maintained%3F-yes-green.svg)  

# Roguelike-tutorial-2020

!["ID: in rainbow text: 'Roguelikedev does the Complete Roguelike Tutorial' and in gray text in the corner: 'v4.0 2020.6-7', over a grayscale ascii dungeon"](https://i.imgur.com/sgsO37A.png)  

## Part 2 - The generic Entity, the render functions, and the map 

Added Camera

![Gif Part 2](https://raw.githubusercontent.com/jolav/roguelike-tutorial/master/assets/gifs/part2.gif)


## Compilation on Linux

>1. Install [Go](https://golang.org/dl/)  
>2. Download [BearLibTerminal](http://foo.wyrd.name/en:bearlibterminal#download)
>3. Put libBearLibTerminal.so into /usr/lib
>4. Move BearLibTerminal.go and BearLibterminal.h into roguelike-tutorial/bearlibterminal  
>5. `go mod init roguelike`. Now you can import library using go modules `import (blt "roguelike/bearlibterminal)`  
>6. `go build -ldflags="-X 'main.releaseDate=$(date -u +%F_%T)'"`  


## Cross Compile on Linux for Windows
>1. `apt install gcc-mingw-w64-x86-64 gcc-multilib`  
>2. Put windows BearLibTerminal.dll and BearLibTerminal.lib into /usr/x86_64-w64-mingw32/lib  
>3. `GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc  go build -ldflags="-X 'main.when=$(date -u +%F_%T)'" -o roguelike.exe`  


## Distribution

#### LINUX

Package in the same folder
- roguelike binary  
- libBearLibTerminal.so  
- start.sh which contains `LD_LIBRARY_PATH="$(dirname "$0")" "$(dirname "$0")/roguelike"`  
- assets folder

enter the folder and execute `./start`
`
#### WINDOWS

Package in the same folder
- roguelike.exe  
- BearLibTerminal.dll
- assets folder

enter the folder and double click roguelike.exe






