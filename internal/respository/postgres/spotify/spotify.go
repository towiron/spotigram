package spotify

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/towiron/spotigram/internal/interfaces"
	"github.com/towiron/spotigram/internal/pkg/config"
	"github.com/towiron/spotigram/internal/pkg/global"
	"github.com/towiron/spotigram/internal/pkg/logger"
	"go.uber.org/fx"
	"log"
	"sync"
	"time"
)

type Repository struct {
	logger logger.Logger
	config config.Configer
	db     interfaces.DB
	client *resty.Client
	token  string
	mu     sync.RWMutex
}

type Options struct {
	fx.In
	fx.Lifecycle
	Logger logger.Logger
	Config config.Configer
	DB     interfaces.DB
}

var Module = fx.Provide(
	fx.Annotate(New, fx.As(new(interfaces.RepositorySpotify))),
)

var _ interfaces.RepositorySpotify = (*Repository)(nil)

func New(opts Options) *Repository {
	client := resty.New().
		SetRetryCount(opts.Config.Int(global.ENV_REQUEST_RETRY_COUNT)).
		SetRetryWaitTime(opts.Config.Duration(global.ENV_REQUEST_RETRY_WAIT_TIME)).
		SetRetryMaxWaitTime(opts.Config.Duration(global.ENV_REQUEST_RETRY_MAX_WAIT_TIME))

	repository := &Repository{
		logger: opts.Logger,
		config: opts.Config,
		db:     opts.DB,
		client: client,
	}

	repository.updateToken()

	opts.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go repository.start()
			return nil
		},
		OnStop: func(context.Context) error { return nil },
	})

	return repository
}

func (r *Repository) start() {
	ticker := time.NewTicker(r.config.Duration(global.ENV_SPOTIFY_TOKEN_LIFE_TIME))

	for range ticker.C {
		r.updateToken()
	}
}

func (r *Repository) updateToken() {
	r.mu.Lock()
	defer r.mu.Unlock()

	formData := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     r.config.String(global.ENV_SPOTIFY_CLIENT_ID),
		"client_secret": r.config.String(global.ENV_SPOTIFY_CLIENT_SECRET),
	}

	resp, err := r.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		Post(r.config.String(global.ENV_SPOTIFY_AUTH_URL))

	if err != nil {
		log.Printf("Failed to update token: %v", err)
		return
	}

	if err := json.Unmarshal(resp.Body(), &tokenResponse); err != nil {
		log.Printf("Failed to parse token response: %v", err)
		return
	}

	r.token = tokenResponse.AccessToken
}
