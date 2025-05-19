package config

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Configer interface {
	Bool(key string) bool
	Float64(key string) float64
	Int(key string) int
	String(key string) string
	Time(key string) time.Time
	Duration(key string) time.Duration
	StringSlice(key string) []string
}

var Module = fx.Provide(New)

func New() Configer {
	cfg := viper.New()

	cfg.AddConfigPath("./configs")
	cfg.SetConfigName(".env")
	cfg.SetConfigType("env")

	err := cfg.ReadInConfig()
	if err != nil {
		panic(err)
	}
	cfg.WatchConfig()

	cfg.AutomaticEnv()

	return &config{cfg: cfg}
}

type config struct {
	cfg *viper.Viper
}

func (c *config) Bool(key string) bool {
	return c.cfg.GetBool(key)
}

func (c *config) Float64(key string) float64 {
	return c.cfg.GetFloat64(key)
}

func (c *config) Int(key string) int {
	return c.cfg.GetInt(key)
}

func (c *config) String(key string) string {
	return c.cfg.GetString(key)
}

func (c *config) Time(key string) time.Time {
	return c.cfg.GetTime(key)
}

func (c *config) Duration(key string) time.Duration {
	return c.cfg.GetDuration(key)
}

func (c *config) StringSlice(key string) []string {
	return c.cfg.GetStringSlice(key)
}
