package main

import p8 "github.com/drpaneas/pigo8"

type game struct{}

func (g *game) Init() {}

func (g *game) Update() {}

func (g *game) Draw() {
    p8.Cls(1)  // Dark blue background
    
    // Draw sprite normally
    p8.Spr(1, 40, 60)
    
    // Draw same sprite with yellow transparent (shows background)
    p8.Palt(10, true)  // Yellow transparent
    p8.Spr(1, 80, 60)
    p8.Palt()  // Reset
}

func main() {
    p8.InsertGame(&game{})
    p8.Play()
}
