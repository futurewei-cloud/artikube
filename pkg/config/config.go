package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

//Config returns the pointer of Viper for service
type Config struct {
	*viper.Viper
}

//NewConfig returns the new configration set with all default values
func NewConfig() *Config {
	conf := &Config{
		Viper: viper.New(),
	}
	conf.SetConfigType("yaml")
	conf.SetDefaults()
	return conf
}

//SetDefaults sets all configuration with the default values
func (conf *Config) SetDefaults() {
	for key, configVar := range configVars {
		conf.SetDefault(key, configVar.Default)
	}
}

//ReadConfigFileFromCLIContext update a config based on flgas set in CLI context
func (conf *Config) ReadConfigFileFromCLIContext(c *cli.Context) error {
	if configFilePath := c.String("config"); configFilePath != "" {
		if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("config files \"%s\" does not exist", configFilePath))
		}

		fileExt := filepath.Ext(configFilePath)
		if fileExt != ".yaml" && fileExt != ".yml" && fileExt != "" {
			return errors.New(fmt.Sprintf("The config file extension is not supported.  Please replace the file with .yaml/.yml"))
		}

		base := strings.TrimSuffix(filepath.Base(configFilePath), fileExt)
		dir := filepath.Dir(configFilePath)
		conf.SetConfigName(base)
		conf.AddConfigPath(dir)
	}

	for key, configVar := range configVars {
		if flag := configVar.CLIFlag; flag != nil {
			name := flag.Names();
			if  c.IsSet(name[0]) {
				switch configVar.Type {
				case stringType:
					conf.Set(key, c.String(name[0]))
				case intType:
					conf.Set(key, c.Int(name[0]))
				case boolType:
					conf.Set(key, c.Bool(name[0]))
				case durationType:
					conf.Set(key, c.Duration(name[0]))
				}
			}
		}
	}

	return nil
}
