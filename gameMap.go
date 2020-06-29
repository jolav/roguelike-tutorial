/* */

package main

func (m *GameMap) fillMapBlockedTiles() {
	m.Tiles = make([][]*Tile, m.Width)
	for i, _ := range m.Tiles {
		m.Tiles[i] = make([]*Tile, m.Height)
	}

	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			m.Tiles[x][y] = &Tile{true, true}
		}
	}
}

func (m *GameMap) IsBlocked(x int, y int) bool {
	if m.Tiles[x][y].Blocked {
		return true
	} else {
		return false
	}
}
