package config

type AppConfig struct {
	HttpPort  int        `yaml:"http_port"`
	HttpProxy string     `yaml:"http_proxy,omitempty"`
	Filter    FilterRule `yaml:"filter"`
}

type FilterRule struct {
	Host []string `yaml:"host"`
}
