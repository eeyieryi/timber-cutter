package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func handleInput(g *Game) (err error) {
	// TODO: REMOVE
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ErrTerminated
	}

	switch g.Scene.Current {
	case NewGameScene:
		if someKeysJustPressed(ebiten.KeySpace) {
			g.Play()
		}
	case PlayingScene:
		if someKeysJustPressed(ebiten.KeyEscape) {
			g.Pause()
			return
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
				for index, sectionPos := range g.State.Tree[:2] {
					switch sectionPos {
					case PosNone:
						continue
					case PosLeft, PosRight:
						if goTo == sectionPos {
							return true, index == 0
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
			g.Play()
		}
	case GameOverScene:
		if someKeysJustPressed(ebiten.KeySpace) {
			g.ResetGame()
		}
	}

	return
}

func someKeysJustPressed(keys ...ebiten.Key) bool {
	for _, k := range keys {
		if inpututil.IsKeyJustPressed(k) {
			return true
		}
	}
	return false
}
