# Complete Roguelike Tutorial from r/roguelikedev

!["ID: in rainbow text: 'Roguelikedev does the Complete Roguelike Tutorial' and in gray text in the corner: 'v4.0 2020.6-7', over a grayscale ascii dungeon"](https://i.imgur.com/sgsO37A.png)  

1. [Go + BearLibTerminal](#go-bearlibterminal)  
2. [Javascript](#javascript)  

# **Go-BearLibTerminal**

## CURRENT STATUS GIF

Part 4 + DIRECTIONAL FOV

![Gif Part 5](https://raw.githubusercontent.com/jolav/roguelike-tutorial/master/Go/assets/gifs/part5.gif)


## CONTROLS  

&uarr; Go Ahead  
&darr; Go Back  
&larr; Turn Left  
&rarr; Turn Right  
ALT + &larr; Strafe Left  
ALT + &rarr; Strafe Right  


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

# **Javascript**  

## **[Play Online](https://roguelike-tutorial-javascript.netlify.app/)**  

## CONTROLS  

NUMPAD  
"1" -> DOWNLEFT  
"2" -> DOWN  
"3" -> DOWNRIGHT  
"4" -> LEFT  
"5" -> SKIP  
"6" -> RIGHT  
"7" -> UPLEFT  
"8" -> UP  
"9" -> UPRIGHT  

KEYBOARD  
"b" -> DOWNLEFT  
"j" -> DOWN  
"n" -> DOWNRIGHT  
"h" -> LEFT  
"t" -> SKIP  
"l" -> RIGHT  
"y" -> UPLEFT  
"k" -> UP  
"u" -> UPRIGHT  

KEYBOARD  
"q" -> LOOT  
"1" -> EAT  
"3" -> HEAL  
"f" -> FIRE  
"r" -> SELECT  

## ROADMAP

- [X] **0.1.5 History minor change**  
- [X] **0.1.4 Refactoring: history has its own class/object**  
- [X] **0.1.3 movement at constant speed keeping key pressed down**  
- [X] **0.1.2 Auto unselect targets**  
- [X] **0.1.1 Colors Palette + History improved + ensure render order + unicode use**  
- [X] **0.1.0 Minimal full playable version**   
- [X] **0.0.14 Leaving The Shelter**
- [X] **0.0.13 Ranged Combat**  
- [X] **0.0.12 Inventory and items**  
- [X] **0.0.11 Creating UI**  
- [X] **0.0.10 Melee**  
- [X] **0.0.9 Turn Time System**  
- [X] **0.0.8 Moving Enemies**  
- [X] **0.0.7 Placing Enemies**  
- [X] **0.0.6 Generating Shelter**  
- [X] **0.0.5 Field of View**  
- [X] **0.0.4 Camera**  
- [X] **0.0.3 Basic Map**   
- [X] **0.0.2 Moving around @**  
- [X] **0.0.1 Setting Up Client-Core Relation**  
- [X] **0.0.0 Start**  

