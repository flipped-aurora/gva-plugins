package config

type Zabbix struct {
	url      string `mapstructure:"url" json:"url" yaml:"url"`
	username string `mapstructure:"username" json:"username" yaml:"username"`
	password string `mapstructure:"password" json:"password" yaml:"password"`
}
