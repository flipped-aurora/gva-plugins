package config

type TencentSMS struct {
	SecretId  string `mapstructure:"secret-id" json:"secret-id" yaml:"secret-id"`
	SecretKey string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
	SdkAppId  string `mapstructure:"sdk-app-id" json:"sdk-app-id" yaml:"sdk-app-id"`
	SignName  string `mapstructure:"sign-name" json:"sign-name" yaml:"sign-name"`
}
