package repository

import (
	"github.com/towiron/spotigram/internal/respository/postgres"
	"github.com/towiron/spotigram/internal/respository/postgres/spotify"
	"go.uber.org/fx"
)

var Module = fx.Options(
	postgres.Module,
	spotify.Module,
)
