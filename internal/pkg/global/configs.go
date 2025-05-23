package global

var (
	ENV_TELEGRAM_BOT_TOKEN = "TELEGRAM_BOT_TOKEN"

	ENV_REPOSITORY_POSTGRES_DSN                     = "REPOSITORY_POSTGRES_DSN"
	ENV_REPOSITORY_POSTGRES_MAX_IDLE_CONNECTIONS    = "REPOSITORY_POSTGRES_MAX_IDLE_CONNECTIONS"
	ENV_REPOSITORY_POSTGRES_MAX_OPEN_CONNECTIONS    = "REPOSITORY_POSTGRES_MAX_OPEN_CONNECTIONS"
	ENV_REPOSITORY_POSTGRES_CONNECTION_MAX_LIFETIME = "REPOSITORY_POSTGRES_CONNECTION_MAX_LIFETIME"

	ENV_REQUEST_RETRY_COUNT         = "REQUEST_RETRY_COUNT"
	ENV_REQUEST_RETRY_WAIT_TIME     = "REQUEST_RETRY_WAIT_TIME"
	ENV_REQUEST_RETRY_MAX_WAIT_TIME = "REQUEST_RETRY_MAX_WAIT_TIME"

	ENV_SPOTIFY_TOKEN_LIFE_TIME = "SPOTIFY_TOKEN_LIFE_TIME"
	ENV_SPOTIFY_AUTH_URL        = "SPOTIFY_AUTH_URL"
	ENV_SPOTIFY_CLIENT_ID       = "SPOTIFY_CLIENT_ID"
	ENV_SPOTIFY_CLIENT_SECRET   = "SPOTIFY_CLIENT_SECRET"
)
