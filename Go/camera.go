/* */

package main

type Camera struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (cam *Camera) move(pj *Entity, m *GameMap) {
	cam.X = pj.X - (cam.Width / 2)
	cam.Y = pj.Y - (cam.Height / 2)
	if cam.X < 0 {
		cam.X = 0
	}
	if cam.Y < 0 {
		cam.Y = 0
	}
	if cam.X > m.Width-cam.Width {
		cam.X = m.Width - cam.Width
	}
	if cam.Y > m.Height-cam.Height {
		cam.Y = m.Height - cam.Height
	}
}

func (cam *Camera) mapToCameraCoords(mapX, mapY int) (camX, camY int) {
	camX = mapX - cam.X
	camY = mapY - cam.Y
	if camX < 0 || camY < 0 || camX >= cam.Width || camY >= cam.Height {
		return -1, -1
	}
	return camX, camY
}
