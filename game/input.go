package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func handleInput(g *Game) error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ErrTerminated
	}

	switch g.Scene.Current {
	case NewGameScene:
		if someKeysJustPressed(ebiten.KeySpace) {
			g.StartNewGame()
		}
	case PlayingScene:
		if someKeysJustPressed(ebiten.KeyEscape) {
			g.Scene.Current = PauseScene
		}

		isLeftPressed := someKeysJustPressed(ebiten.KeyArrowLeft, ebiten.KeyH)
		isRightPressed := someKeysJustPressed(ebiten.KeyArrowRight, ebiten.KeyL)

		goTo := PosNone
		if isLeftPressed && !isRightPressed {
			goTo = PosLeft
		} else if isRightPressed && !isLeftPressed {
			goTo = PosRight
		}

		switch goTo {
		case PosLeft, PosRight:
			hit, root := func() (hit, root bool) {
				for i, sectionPos := range g.State.Tree[:2] {
					switch sectionPos {
					case PosNone:
						continue
					case PosLeft, PosRight:
						if goTo == sectionPos {
							return true, i == 0
						}
					}
				}
				return
			}()
			g.State.PlayerPos = goTo
			if hit {
				g.Scene.Current = GameOverScene
			}
			if !root {
				g.State.UpdateTree()
			}
		}
	case PauseScene:
		if someKeysJustPressed(ebiten.KeyEscape) {
			g.Scene.Current = PlayingScene
		}
	case GameOverScene:
		if someKeysJustPressed(ebiten.KeySpace) {
			g.StartNewGame()
		}
	}

	return nil
}

func someKeysJustPressed(keys ...ebiten.Key) bool {
	for _, k := range keys {
		if inpututil.IsKeyJustPressed(k) {
			return true
		}
	}
	return false
}
