package views

import (
	"github.com/dradtke/allegory"
	"github.com/dradtke/allegory-example/main/game/commands"
	"github.com/dradtke/allegory-example/main/game/field"
	"github.com/dradtke/go-allegro/allegro"
)

type HumanView struct {
	allegory.BaseView

	leftDown  bool
	rightDown bool
	upDown    bool
	downDown  bool
}

var _ allegory.View = (*HumanView)(nil)

func (v *HumanView) HandleEvent(msg interface{}) bool {
	switch e := msg.(type) {
	case allegro.KeyDownEvent:
		switch e.KeyCode() {
		case allegro.KEY_LEFT:
			v.leftDown = true
			return true
		case allegro.KEY_RIGHT:
			v.rightDown = true
			return true
		case allegro.KEY_UP:
			v.upDown = true
			return true
		case allegro.KEY_DOWN:
			v.downDown = true
			return true
		}

	case allegro.KeyUpEvent:
		switch e.KeyCode() {
		case allegro.KEY_LEFT:
			v.leftDown = false
			return true
		case allegro.KEY_RIGHT:
			v.rightDown = false
			return true
		case allegro.KEY_UP:
			v.upDown = false
			return true
		case allegro.KEY_DOWN:
			v.downDown = false
			return true
		}
	}
	return false
}

func (v *HumanView) UpdateView() {
	if v.rightDown && !v.leftDown {
		field.Player.HandleCommand(commands.PlayerMoveRight)
	} else if !v.rightDown && v.leftDown {
		field.Player.HandleCommand(commands.PlayerMoveLeft)
	}
}
