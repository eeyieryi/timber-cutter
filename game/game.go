package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	Title        = "Timber Cutter"
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	State *GameState
	Scene GameScene
}

func (g *Game) ResetGame() {
	g.State = NewGameState()
	g.Scene.Current = NewGameScene
}

func (g *Game) Play() {
	g.Scene.Current = PlayingScene
}

func (g *Game) Pause() {
	g.Scene.Current = PauseScene
}

func (g *Game) Update() (err error) {
	switch g.State {
	case nil:
		g.ResetGame()
	}
	err = handleInput(g)
	return
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
