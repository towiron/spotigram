package main

import (
	"github.com/towiron/spotigram/internal/app"
	"github.com/towiron/spotigram/internal/pkg"
)

func main() {
	app.New(
		pkg.Module,
	).Run()
}
