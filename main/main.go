package main

import (
	"github.com/dradtke/gopher"
	"github.com/dradtke/gopher-example/main/loading"
	"github.com/dradtke/gopher/config"
)

func main() {
	config.SetDisplaySize(700, 350)
	config.SetPackageRoot("github.com/dradtke/gopher-example")

	gopher.Init(new(loading.LoadingState))
	defer gopher.Cleanup()
	gopher.Loop()
}
