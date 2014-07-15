package actors

import (
	"fmt"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/gopher"
	"github.com/dradtke/gopher-example/main/game/globals"
	"github.com/dradtke/gopher/cache"
)

type Box struct {
	gopher.BaseActor
	img *allegro.Bitmap
}

type BoxState int

const (
	_ BoxState = iota
	NormalBox
	AltBox
	EmptyBox
	WarningBox
	ExplosiveBox
)

func NewBox(x, y float32, state BoxState) *Box {
	var img string
	switch state {
	case NormalBox:
		img = "box.png"
	case AltBox:
		img = "boxAlt.png"
	case EmptyBox:
		img = "boxEmpty.png"
	case WarningBox:
		img = "boxWarning.png"
	case ExplosiveBox:
		img = "boxExplosiveAlt.png"
	default:
		panic(fmt.Sprintf("unknown box state: %d", state))
	}
	return &Box{BaseActor: gopher.BaseActor{x * globals.TILE_SIZE, y * globals.TILE_SIZE, 0, 0}, img: cache.Image(globals.TILE_ASSETS + img)}
}

func (b *Box) UpdateActor() {
	// Move?
}

func (b *Box) RenderActor(delta float32) {
	x, y := b.CalculatePos(delta)
	b.img.Draw(x, y, allegro.FLIP_NONE)
}
