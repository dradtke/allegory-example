package loading

import (
	"github.com/dradtke/allegory"
	"github.com/dradtke/allegory-example/main/game/globals"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/font"
	"os"
	"path/filepath"
)

type LoadingState struct {
	allegory.BaseState
}

func AddImages(images *[]string, dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".png" {
			*images = append(*images, path)
		}
		return nil
	})
}

func (s *LoadingState) InitState() {
	go func() {
		images := make([]string, 0)
		AddImages(&images, globals.TILE_ASSETS)
		AddImages(&images, globals.PLAYER_ASSETS)
		allegory.RunProcess(&load{Images: images})
	}()
}

func (s *LoadingState) RenderState(delta float32) {
	font.DrawText(allegory.BuiltinFont(), allegro.MapRGB(255, 255, 255), 10, 10, font.ALIGN_LEFT, "Loading...")
}
