package configuration

type Config struct {
	API struct {
		URL string `yaml:"URL", envconfig:"API_URL"`
	} `yaml:"API"`
}
