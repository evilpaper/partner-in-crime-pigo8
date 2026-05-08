package main

import p8 "github.com/drpaneas/pigo8"

type Game struct{}

func (g *Game) Init() {}

func (g *Game) Update() {}

func (g *Game) Draw() {
  p8.Cls(0)
}

func main() {
    p8.InsertGame(&Game{})
    p8.Play()
}
