package main

import p8 "github.com/drpaneas/pigo8"

type game struct{
    x, y float64
}

func (g *game) Init()   {
    g.x = 60
    g.y = 60
}
func (g *game) Update() {
    if p8.Btn(p8.LEFT) { g.x-- }
    if p8.Btn(p8.RIGHT) { g.x++ }
    if p8.Btn(p8.UP) { g.y-- }
    if p8.Btn(p8.DOWN) { g.y++ }
}

func (g *game) Draw() {
    p8.Cls(0)
    p8.Rectfill(g.x, g.y, g.x+8, g.y+8, 8) // Red square
}


func main() {
    p8.InsertGame(&game{})
    p8.Play()
}