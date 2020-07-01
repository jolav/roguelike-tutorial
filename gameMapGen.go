/* */

package main

type GameMap struct {
	Width  int
	Height int
	Tiles  [][]*Tile
}

type Wall struct {
	X   int
	Y   int
	Nei int
	Dir string
}

type Feature struct {
	Width  int
	Height int
}

type Room struct {
	X      int
	Y      int
	Width  int
	Height int
}

const (
	ROOM_TRIES          = 1000
	ROOMS               = 20
	MIN_SIZE_ROOM       = 6
	MAX_SIZE_ROOM       = 15
	MIN_LENGTH_CORRIDOR = 2
	MAX_LENGTH_CORRIDOR = 10
	CORRIDOR_PERCENT    = 40
)

func (m *GameMap) initializeMap() {
	m.fillMapBlockedTiles()
	m.createSingleRoomInCenter()
	success := 0
	for tries := 1; tries < ROOM_TRIES; tries++ {
		w := m.pickRandomWallFromAnyRoom()
		//prettyPrintStruct(w)
		f := m.pickRandomFeature()
		//prettyPrintStruct(f)
		r := m.convertFeatureToRoom(w, f)
		//prettyPrintStruct(r)
		if m.checkIsRoomForFeature(r) {
			m.fillRectRoom(r)
			m.fillWall(w)
			success++
			if success >= ROOMS {
				return
			}
		} else {
			//m.Tiles[w.X][w.Y] = &Tile{true, true} // clean randomWall
		}

	}
}

func (m *GameMap) pickRandomWallFromAnyRoom() (w *Wall) {
	w = new(Wall)
	var found bool = false
	var limit int = 0
	for !found && limit < 5000 {
		w.X = randomInt(0, m.Width-1)
		w.Y = randomInt(0, m.Height-1)
		if m.Tiles[w.X][w.Y].Blocked && m.notInTheBoardEdge(w.X, w.Y) {
			w.Nei, w.Dir = m.getClearNeighbours(w.X, w.Y)
			if w.Nei == 1 {
				found = true
			}
		}
		limit++
	}
	if found {
		return w
	}
	return &Wall{0, 0, 0, ""}
}

func (m *GameMap) getClearNeighbours(x, y int) (int, string) {
	var nei int = 0
	var dir = "Zero"
	if !m.Tiles[x][y-1].Blocked {
		nei++
		dir = "S"
	}
	if !m.Tiles[x][y+1].Blocked {
		nei++
		dir = "N"
	}
	if !m.Tiles[x+1][y].Blocked {
		nei++
		dir = "W"
	}
	if !m.Tiles[x-1][y].Blocked {
		nei++
		dir = "E"
	}
	return nei, dir
}

func (m *GameMap) notInTheBoardEdge(x, y int) bool {
	if x < 1 || y < 1 || x > m.Width-2 || y > m.Height-2 {
		return false
	}
	return true
}

func (m *GameMap) pickRandomFeature() (f *Feature) {
	f = new(Feature)
	random := randomInt(1, 100)
	switch {
	case random < CORRIDOR_PERCENT:
		f.Width = 1
		f.Height = randomInt(MIN_LENGTH_CORRIDOR, MAX_LENGTH_CORRIDOR)
	case random >= CORRIDOR_PERCENT:
		f.Width = randomInt(MIN_SIZE_ROOM, MAX_SIZE_ROOM)
		f.Height = randomInt(MIN_SIZE_ROOM, MAX_SIZE_ROOM)
	}
	return
}

func (m *GameMap) convertFeatureToRoom(w *Wall, f *Feature) (r *Room) {
	r = new(Room)
	switch {
	case w.Dir == "N":
		r.X = w.X - f.Width/2
		r.Y = w.Y - f.Height
		r.Width = f.Width
		r.Height = f.Height
	case w.Dir == "S":
		r.X = w.X - f.Width/2
		r.Y = w.Y + 1
		r.Width = f.Width
		r.Height = f.Height
	case w.Dir == "E":
		r.X = w.X + 1
		r.Y = w.Y - f.Width/2
		r.Width = f.Height
		r.Height = f.Width
	case w.Dir == "W":
		r.X = w.X - f.Height
		r.Y = w.Y - f.Width/2
		r.Width = f.Height
		r.Height = f.Width
	}
	return r

}

func (m *GameMap) checkIsRoomForFeature(r *Room) bool {
	if r.X+r.Width > m.Width-1 || r.Y+r.Height > m.Height-1 {
		return false
	}
	if r.X <= 0 || r.Y <= 0 { // =0 avoid rooms just in the edge
		return false
	}
	originX := r.X // (m.Width - r.Width) / 2
	originY := r.Y // (m.Height - r.Height) / 2
	for x := 0; x < r.Width; x++ {
		for y := 0; y < r.Height; y++ {
			if !m.Tiles[originX+x][originY+y].Blocked {
				return false
			}
		}
	}
	return true
}

func (m *GameMap) createSingleRoomInCenter() {
	width := randomInt(MIN_SIZE_ROOM, MAX_SIZE_ROOM)
	height := randomInt(MIN_SIZE_ROOM, MAX_SIZE_ROOM)
	r := &Room{
		X:      (m.Width - width) / 2,
		Y:      (m.Height - height) / 2,
		Width:  width,
		Height: height,
	}
	m.fillRectRoom(r)
}

func (m *GameMap) fillWall(w *Wall) {
	m.Tiles[w.X][w.Y] = &Tile{false, false, false, false}
}

func (m *GameMap) fillRectRoom(r *Room) {
	originX := r.X
	originY := r.Y
	for x := 0; x < r.Width; x++ {
		for y := 0; y < r.Height; y++ {
			m.Tiles[originX+x][originY+y] = &Tile{false, false, false, false}
		}
	}
}
