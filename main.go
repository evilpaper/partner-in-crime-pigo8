package main

import p8 "github.com/drpaneas/pigo8"

type Paddle struct {
    x, y, width, height, speed float64
    color int
}

type Ball struct{
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

func (g *Game) Init() {

    // Initialize positions

    // Player paddle on the left
    g.player = Paddle{
        x: 4, y: 54,
        width: 4, height: 20,
        speed: 4, color: 12,
    }

    // Computer paddle on the right
    g.computer = Paddle{
        x: 120, y: 54,
        width: 4, height: 20,
        speed: 1.5, color: 8,
    }

    // Ball in the center
    g.ball = Ball{
        x: 64, y: 32, size: 4,
        dx: 2, dy: 1,
        color: 7,
    }
}

func collides(ball Ball, paddle Paddle) bool {
    return ball.x+ball.size >= paddle.x &&
           ball.x <= paddle.x+paddle.width &&
           ball.y+ball.size >= paddle.y &&
           ball.y <= paddle.y+paddle.height
}

func (g *Game) Update() {

    // Player movement
    if p8.Btn(p8.UP) && g.player.y > 0 {
        g.player.y -= g.player.speed
    }
    if p8.Btn(p8.DOWN) && g.player.y + g.player.height < 128 {
        g.player.y += g.player.speed
    }

    // Move the ball
    g.ball.x += g.ball.dx
    g.ball.y += g.ball.dy

    // Bounce off walls
    if g.ball.y <= 0 || g.ball.y+g.ball.size >= 128 {
        g.ball.dy = -g.ball.dy
    }

    // Paddle collision
    if collides(g.ball, g.player) || collides(g.ball, g.computer) {
        g.ball.dx = -g.ball.dx
    }

    // Score when ball exits
    if g.ball.x < 0 {
        g.computerScore++
        g.resetBall()
    }
    if g.ball.x > 128 {
        g.playerScore++
        g.resetBall()
    }
}

func (g *Game) resetBall() {
    g.ball.x = 62
    g.ball.y = 62
    g.ball.dx = -g.ball.dx  // Serve toward last scorer
}

func (g *Game) Draw() {
  p8.Cls(0)
}

func main() {
    p8.InsertGame(&Game{})
    p8.Play()
}
