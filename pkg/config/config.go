package config

import (
	"github.com/spf13/viper"
)

//Configuration for service
type Config struct {
	*viper.Viper
}

//Return the new configration set with all default values
func NewConfig() *Config {
	conf := &Config{
		Viper: viper.New(),
	}
	conf.SetConfigType("yaml")
	conf.SetDefaults()
	return conf
}

//Set all configuration with the default values
func (conf *Config) SetDefaults() {
	for key, configVar := range configVars {
		conf.SetDefault(key, configVar.Default)
	}
}
