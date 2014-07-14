package commands

type PlayerCommand int

const (
	_ PlayerCommand = iota
	PlayerMoveRight
	PlayerMoveLeft
)
