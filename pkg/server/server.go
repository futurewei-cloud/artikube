package artikube

import (
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
		//Logger
		//router
	}

	//ServerOptions are options
	ServerOptions struct {
		StorageBackend storage.Backend
		LogJSON        bool
		Debug          bool
		ArtifactURL    string
	}
)

func NewServer(options ServerOptions) *Server {
	var artifactURL string
	if options.ArtifactURL != "" {
		artifactURL = options.ArtifactURL
	}

	server := &Server{
		StorageBackend: options.StorageBackend,
		ArtifactURL:    artifactURL,
	}

	return server
}

func (server *Server) Listen(port int) {
	//server.Router.start(port)
}
