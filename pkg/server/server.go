package artikube

import (
	arti_logger "github.com/futurewei-cloud/artikube/pkg/server/logger"
	arti_router "github.com/futurewei-cloud/artikube/pkg/server/router"
	"github.com/futurewei-cloud/artikube/pkg/storage"
	"github.com/gin-gonic/gin"
)

type (
	//Router handles all incoming http requests
	Router struct {
		*gin.Engine
	}

	Server struct {
		StorageBackend storage.Backend
		ArtifactURL    string
		Logger         *arti_logger.Logger
		Router         *arti_router.Router
	}

	//ServerOptions are options
	ServerOptions struct {
		StorageBackend storage.Backend
		Logger         *arti_logger.Logger
		Router         *arti_router.Router
		LogJSON        bool
		Debug          bool
		ArtifactURL    string
	}
)

func NewServer(options ServerOptions) (*Server, error) {
	//logger := arti_logger.NewLogger(&arti_logger.LoggerConfiguration{})
	var artifactURL string
	if options.ArtifactURL != "" {
		artifactURL = options.ArtifactURL
	}

	server := &Server{
		StorageBackend: options.StorageBackend,
		Logger:         options.Logger,
		Router:         options.Router,
		ArtifactURL:    artifactURL, 
	}

	server.Router.SetRoutes(server.Routes())

	return server, nil
}

func (server *Server) Listen(port int) {

}
