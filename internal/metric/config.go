package metric

type Config struct {
	Enable bool   `koanf:"enable"`
	Host   string `koanf:"host"`
}
