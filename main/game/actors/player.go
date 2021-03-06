package actors

import (
	"github.com/dradtke/allegory"
	"github.com/dradtke/allegory-example/main/game/commands"
	"github.com/dradtke/allegory-example/main/game/globals"
	"github.com/dradtke/allegory/cache"
	"github.com/dradtke/go-allegro/allegro"
)

type Player struct {
	allegory.BaseActor

	img     *allegro.Bitmap
	flipped bool
	speed   float32
}

func (p *Player) InitActor() {
	p.img = cache.Image(globals.PLAYER_ASSETS + "p1_stand.png")
	p.speed = 4
}

func (p *Player) UpdateActor() {
}

func (p *Player) RenderActor(delta float32) {
	x, y := p.CalculatePos(delta)
	flags := allegro.FLIP_NONE
	if p.flipped {
		flags = allegro.FLIP_HORIZONTAL
	}
	p.img.Draw(x, y, flags)
}

func (p *Player) HandleCommand(cmd interface{}) {
	switch cmd {
	case commands.PlayerMoveRight:
		p.Move(p.speed, 0)
		p.flipped = false

	case commands.PlayerMoveLeft:
		p.Move(-p.speed, 0)
		p.flipped = true
	}
}
