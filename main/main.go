package main

import (
	"github.com/dradtke/allegory"
	"github.com/dradtke/allegory-example/main/loading"
	"github.com/dradtke/allegory/config"
)

func main() {
	config.SetDisplaySize(700, 350)
	config.SetPackageRoot("github.com/dradtke/allegory-example")

	allegory.Init(new(loading.LoadingState))
	defer allegory.Cleanup()
	allegory.Loop()
}
