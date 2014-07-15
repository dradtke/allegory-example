package globals

import (
	"path/filepath"
)

const (
	ASSETS_DIR = "assets"
)

var (
	PLAYER_ASSETS = filepath.Join(ASSETS_DIR, "player") + string(filepath.Separator)
	TILE_ASSETS   = filepath.Join(ASSETS_DIR, "tiles") + string(filepath.Separator)
)
