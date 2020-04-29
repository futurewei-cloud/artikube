package config

import "github.com/urfave/cli"

type configVar struct {
	Type    configVarType
	Default interface{}
	CLIFlag cli.Flag
}

type configVarType string

var CLIFlags []cli.Flag

var (
	stringType   configVarType = "string"
	intType      configVarType = "int"
	boolType     configVarType = "bool"
	durationType configVarType = "time.Duration"
)

var configVars = map[string]configVar{
	"debug": {
		Type:    boolType,
		Default: false,
		CLIFlag: &cli.BoolFlag{
			Name:  "debug",
			Usage: "display debug message",
			//EnvVars: "DEBUG",
		},
	},
}
