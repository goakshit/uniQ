package config

type connConfig struct {
	ConnHost string
	ConnPort int
	ConnType string
}

// config - Application configuration
type config struct {
	ConnConfig connConfig
}
