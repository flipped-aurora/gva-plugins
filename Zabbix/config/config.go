package config

type Server struct {
	Zabbix Zabbix `mapstructure:"zabbix" json:"zabbix" yaml:"zabbix"`
}
