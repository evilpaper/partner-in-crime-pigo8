package main

import p8 "github.com/drpaneas/pigo8"

type game struct {
    stars [][2]int  // x, y pairs
}

func (g *game) Init() {
    // Create 50 random stars
    for i := 0; i < 50; i++ {
        x := p8.Rnd(128)
        y := p8.Rnd(128)
        g.stars = append(g.stars, [2]int{x, y})
    }
}

func (g *game) Update() {
    // Update the stars
    for _, star := range g.stars {
        star[1]++
        if star[1] > 128 {
            star[1] = 0
        }
    }
}

func (g *game) Draw() {
    p8.Cls(0)
    for _, star := range g.stars {
        // Random twinkle: white or light gray
        color := 7
        if p8.Rnd(10) < 3 {
            color = 6
        }
        p8.Pset(star[0], star[1], color)
    }
}

func main() {
    p8.InsertGame(&game{})
    p8.Play()
}
