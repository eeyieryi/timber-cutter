package game

import (
	"bytes"
	"image"

	"github.com/eeyieryi/timber-cutter/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	blockImg *ebiten.Image
)

type Scene int

const (
	NewGameScene Scene = iota
	PlayingScene
	PauseScene
	GameOverScene
)

type GameScene struct {
	Current Scene
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Img_Block_png))
	if err != nil {
		panic(err)
	}
	blockImg = ebiten.NewImageFromImage(img)
}

func (gs *GameScene) Draw(game *Game, screen *ebiten.Image) {
	switch gs.Current {
	case NewGameScene:
		drawTree(screen, game.State)
		drawPlayer(screen, game.State)
		// TODO: drawNewGameOverlay()
	case PlayingScene:
		drawPlayer(screen, game.State)
		drawTree(screen, game.State)
		// TODO: drawScore()
	case PauseScene:
		drawTree(screen, game.State)
		// TODO: drawPauseOverlay()
		// TODO: drawScore()
	case GameOverScene:
		drawPlayer(screen, game.State)
		drawTree(screen, game.State)
		// TODO: drawScore()
		// TODO: drawGameOverOverlay()
	}
}

const (
	treeWidth    float32 = 64
	treeDrawGap  int     = 5
	playerWidth  float32 = 32
	playerHeight float32 = 64
	branchWidth  float32 = 64
	branchHeight float32 = 16
	gap          float32 = playerHeight
)

func getScreenSize(screen *ebiten.Image) (x, y float32) {
	return float32(screen.Bounds().Dx()), float32(screen.Bounds().Dy())
}

func drawPlayer(screen *ebiten.Image, gameState *GameState) {
	screenW, screenH := getScreenSize(screen)
	var centerX float32 = screenW / 2

	var rectX float32
	switch gameState.PlayerPos {
	case PosLeft:
		leftX := centerX - playerWidth*1.5 - treeWidth/2
		rectX = leftX
	case PosRight:
		rightX := centerX + playerWidth*0.5 + treeWidth/2
		rectX = rightX
	}

	var rectY float32 = screenH - playerHeight

	vector.DrawFilledRect(screen, rectX, rectY, playerWidth, playerHeight, RedColor, false)
}

func drawTree(screen *ebiten.Image, gameState *GameState) {
	screenW, screenH := getScreenSize(screen)
	var centerX float32 = screenW / 2

	for y := 0; y < int(screenH); y += treeDrawGap {
		geoM := ebiten.GeoM{}
		geoM.Scale(4, 4)
		geoM.Translate(float64(centerX-treeWidth/2), float64(y))
		screen.DrawImage(blockImg, &ebiten.DrawImageOptions{GeoM: geoM})
	}

	// vector.DrawFilledRect(screen, centerX, 0, treeWidth, screenH, color.White, false)
	for index, sectionPos := range gameState.Tree {
		if sectionPos == PosNone {
			continue
		}

		var branchX float32
		switch sectionPos {
		case PosLeft:
			branchX = centerX - branchWidth - treeWidth/2
		case PosRight:
			branchX = centerX + treeWidth/2
		}

		var branchY float32 = screenH - branchHeight - float32(index)*gap
		if index == 0 {
			branchY -= branchHeight
		}

		vector.DrawFilledRect(screen, branchX, branchY, branchWidth, branchHeight, GreenColor, false)
	}
}
