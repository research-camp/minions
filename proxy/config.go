package proxy

type Config struct {
	Address string `koanf:"address"`
	BaseURL string `koanf:"base_url"`
}
