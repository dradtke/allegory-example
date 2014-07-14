package main

import (
	"./loading"
	"github.com/dradtke/gopher"
	"github.com/dradtke/gopher/config"
)

func main() {
	config.SetDisplaySize(700, 350)
	gopher.Init(new(loading.LoadingState))
	defer gopher.Cleanup()
	gopher.Loop()
}
