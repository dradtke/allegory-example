package game

import (
	"github.com/dradtke/allegory"
	"github.com/dradtke/allegory-example/main/game/actors"
	"github.com/dradtke/allegory-example/main/game/field"
	"github.com/dradtke/allegory-example/main/game/views"
	"github.com/dradtke/allegory/config"
)

const (
	TILE_SIZE = 70
)

func Play() {
	allegory.NewState(&PlayingState{}, new(views.HumanView))
}

type PlayingState struct {
	allegory.BaseState
}

func (s *PlayingState) InitState() {
	allegory.AddActor(actors.NewBox(0, 4, actors.NormalBox))
	allegory.AddActor(actors.NewBox(1, 4, actors.ExplosiveBox))
	allegory.AddActor(actors.NewBox(2, 4, actors.AltBox))
	allegory.AddActor(actors.NewBox(3, 4, actors.NormalBox))
	allegory.AddActor(actors.NewBox(4, 4, actors.EmptyBox))
	allegory.AddActor(actors.NewBox(5, 4, actors.NormalBox))
	allegory.AddActor(actors.NewBox(6, 4, actors.EmptyBox))
	allegory.AddActor(actors.NewBox(7, 4, actors.NormalBox))
	allegory.AddActor(actors.NewBox(8, 4, actors.WarningBox))
	allegory.AddActor(actors.NewBox(9, 4, actors.NormalBox))

	_, sh := config.DisplaySize()
	field.Player = new(actors.Player)
	field.Player.X = 100
	field.Player.Y = float32(sh - (TILE_SIZE * 2.5))
	allegory.AddActor(field.Player)
}
