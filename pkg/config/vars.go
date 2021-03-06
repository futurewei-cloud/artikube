package config

import "github.com/urfave/cli/v2"

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
			Name:    "debug",
			Usage:   "display debug message",
			EnvVars: []string{"DEBUG"},
		},
	},
	"port": {
		Type:    intType,
		Default: 8080,
		CLIFlag: &cli.IntFlag{
			Name:    "port",
			Usage:   "port to listen on",
			EnvVars: []string{"PORT"},
		},
	},
	"storage.backend": {
		Type:    stringType,
		Default: "./storage",
		CLIFlag: &cli.StringFlag{
			Name:    "storage",
			Usage:   "storage backend, can be one of: local, amazon, google, oracle",
			EnvVars: []string{"STORAGE"},
		},
	},
	"storage.filesystem.rootdir": {
		Type:    stringType,
		Default: "./.store",
		CLIFlag: &cli.StringFlag{
			Name:    "storage-filesystem-rootdir",
			Usage:   "directory to store charts for file system storage backend",
			EnvVars: []string{"STORAGE_STORAGE_ROOTDIR"},
		},
	},
}

func populateCLIFlags() {
	CLIFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config, c",
			Usage:   "artipie configuration file",
			EnvVars: []string{"CONFIG"},
		},
	}
	for _, configVar := range configVars {
		if flag := configVar.CLIFlag; flag != nil {
			CLIFlags = append(CLIFlags, flag)
		}
	}
}

func init() {
	populateCLIFlags()
}
