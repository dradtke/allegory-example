package loading

import (
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/font"
	"github.com/dradtke/gopher"
	"github.com/dradtke/gopher-example/main/game/globals"
	"os"
	"path/filepath"
)

type LoadingState struct {
	gopher.BaseState
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
		gopher.RunProcess(&load{Images: images})
	}()
}

func (s *LoadingState) RenderState(delta float32) {
	font.DrawText(gopher.BuiltinFont(), allegro.MapRGB(255, 255, 255), 10, 10, font.ALIGN_LEFT, "Loading...")
}
