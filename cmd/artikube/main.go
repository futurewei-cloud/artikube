package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/futurewei-cloud/artikube/pkg/config"
	artikube "github.com/futurewei-cloud/artikube/pkg/server"

	//arti_logger "github.com/futurewei-cloud/artikube/pkg/server/logger"
	"github.com/futurewei-cloud/artikube/pkg/storage"
	"github.com/urfave/cli/v2"
)

var (
	crash = log.Fatal

	newServer = artikube.NewServer

	Version string

	Revision string
)

func main() {
	app := &cli.App{
		Name:    "artikube",
		Version: fmt.Sprint("%s (build %s)", Version, Revision),
		Usage:   "Artifact Repository with the support for Maven, NPM, Docker, etc",
		Action:  cliAction,
		Flags:   config.CLIFlags,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Artikube does not start due to the error: %s", err)
	}
}

func cliAction(c *cli.Context) error {
	conf := config.NewConfig()

	backend := backendFromConfig(conf)
	//cache store := storeFromConfig(conf)

	options := artikube.ServerOptions{
		StorageBackend: backend,
		ArtifactURL:    conf.GetString("artifacturl"),
		Debug:          conf.GetBool("debug"),
	}

	server, err := newServer(options)
	server.Listen(conf.GetInt("port"))

	return err
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
