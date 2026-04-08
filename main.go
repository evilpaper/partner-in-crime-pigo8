package main

import p8 "github.com/drpaneas/pigo8"

type game struct{}

func (g *game) Init()   {}
func (g *game) Update() {}
func (g *game) Draw()   { p8.Cls(1) }

func main() {
    p8.InsertGame(&game{})
    p8.Play()
}