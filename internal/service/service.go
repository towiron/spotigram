package service

import (
	"github.com/towiron/spotigram/internal/service/syncer"
	"go.uber.org/fx"
)

var Module = fx.Options(
	syncer.Module,
)
