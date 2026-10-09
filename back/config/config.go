package config

type Config struct {
	PORT string
}

func InitConfig() *Config {
	return &Config{
		PORT: "localhost:4000",
	}
}

// Context keys used for storing request-scoped user information.
const (
	UserIDKey   = "user_id"
	UserNameKey = "user_name"
	RoleKey     = "role"
	JwtSecret   = "change-me-in-production"
)
