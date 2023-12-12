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

func (g *Game) StartNewGame() {
	g.State = NewGameState()
	g.Scene.Current = PlayingScene
}

func (g *Game) Update() (err error) {
	// TODO: Temporary
	if g.Scene.Current == NewGameScene {
		g.StartNewGame()
		return
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
