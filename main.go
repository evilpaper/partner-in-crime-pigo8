package main

import p8 "github.com/drpaneas/pigo8"
import "fmt"

type game struct {
    score int
    lives int
}

func (g *game) Init() {
    g.score = 0
    g.lives = 3
}
func (g *game) Update() {}

func (g *game) Draw() {
    p8.Cls(0)
    
    // Title
    p8.Print("SPACE SHOOTER", 30, 5, 7)
    
    // Score (right-aligned)
    scoreText := fmt.Sprintf("%05d", g.score)
    p8.Print("SCORE:", 85, 5, 6)
    p8.Print(scoreText, 105, 5, 11)
    
    // Lives
    p8.Print("LIVES:", 5, 5, 6)
    for i := 0; i < g.lives; i++ {
        p8.Spr(1, 30+i*10, 3)  // Heart sprites
    }
}

func main() {
    p8.InsertGame(&game{})
    p8.Play()
}
