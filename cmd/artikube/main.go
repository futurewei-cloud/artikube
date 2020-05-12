package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/futurewei-cloud/artikube/pkg/config"
	artikube "github.com/futurewei-cloud/artikube/pkg/server"
	"github.com/futurewei-cloud/artikube/pkg/storage"
	"github.com/urfave/cli"
)

var (
	crash = log.Fatal

	newServer = artikube.NewServer

	Version string

	Revision string
)

func main() {
	app := cli.NewApp()
	app.Name = "artikube"
	app.Version = fmt.Sprint("%s (build %s)", Version, Revision)
	app.Usage = "Artifact Repository with the support for Maven, NPM, Docker, etc"
	//app.Action = cliHandler
	app.Flags = config.CLIFlags
	app.Run(os.Args)
}

func cliHandler(c *cli.Context) {
	conf := config.NewConfig()

	backend := backendFromConfig(conf)
	//cache store := storeFromConfig(conf)

	options := artikube.ServerOptions{
		StorageBackend: backend,
		ArtifactURL:    conf.GetString("artifacturl"),
		Debug:          conf.GetBool("debug"),
	}

	server, err := newServer(options)

	if err != nil {
		crash(err)
	}

	server.Listen(conf.GetInt("port"))
}

func backendFromConfig(conf *config.Config) storage.Backend {
	//Todo: crashIfConfigMissingVars()

	var backend storage.Backend

	storageType := strings.ToLower(conf.GetString("storage.backend"))
	switch storageType {
	case "filesystem":
		backend = filesystemBackendFromConfig(conf)
	// will add more support to different storage type
	default:
		crash("Currently do not support this type of storage", storageType)
	}
	return backend
}

func filesystemBackendFromConfig(conf *config.Config) storage.Backend {
	return storage.Backend(storage.NewFilesystemBackend(
		conf.GetString("storage.filesystem.rootdir"),
	))
}
