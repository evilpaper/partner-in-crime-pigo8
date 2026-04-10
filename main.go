package main

import p8 "github.com/drpaneas/pigo8"

type game struct{}

func (g *game) Init()   {}
func (g *game) Update() {}

func (g *game) Draw() {
    p8.Cls(0)
    p8.Print("game boy!", 52, 70, 3)
}

func main() {
    settings := p8.NewSettings()
    settings.ScreenWidth = 160
    settings.ScreenHeight = 144
    settings.WindowTitle = "Game Boy Style"
    settings.ScaleFactor = 4
    
    p8.InsertGame(&game{})
    p8.PlayGameWith(settings)
}
