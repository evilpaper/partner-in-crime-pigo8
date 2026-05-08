package main

import p8 "github.com/drpaneas/pigo8"

type Paddle struct {
    x, y, width, height, speed float64
    color int
}

type Ball stuct {
    x,  y, size float64
    dx, dy  float64
    color int
}

type Game struct{
    player Paddle
    computer Paddle 
    ball Ball
    playerScore, computerScore int
}

func (g *Game) Init() {}

func (g *Game) Update() {}

func (g *Game) Draw() {
  p8.Cls(0)
}

func main() {
    p8.InsertGame(&Game{})
    p8.Play()
}
