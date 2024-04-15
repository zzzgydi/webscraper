package config

type AppConfig struct {
	HttpPort  int            `yaml:"http_port"`
	HttpProxy string         `yaml:"http_proxy,omitempty"`
	Filter    FilterRule     `yaml:"filter"`
	Chrome    ChromeSettings `yaml:"chrome"`
}

type FilterRule struct {
	Host        []string `yaml:"host"`
	BlockString []string `yaml:"block_string"`
}

type ChromeSettings struct {
	RemoteUrl string `yaml:"remote_url,omitempty"`
	ExecPath  string `yaml:"exec_path,omitempty"`
}
