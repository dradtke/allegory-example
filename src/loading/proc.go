package loading

import (
	"../game"
	"fmt"
	"github.com/dradtke/gopher"
	"github.com/dradtke/gopher/cache"
	"os"
)

type load struct {
	Images  []string
	loading bool
}

var _ gopher.Process = (*load)(nil)

func (p *load) InitProcess() {
	p.loading = true
	go func() {
		errs := cache.LoadImages(p.Images)
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
		p.loading = false
	}()
}

func (p *load) HandleMessage(msg interface{}) error {
	return nil
}

func (p *load) Tick() (bool, error) {
	return p.loading, nil
}

func (p *load) CleanupProcess() {
	game.Play()
}
