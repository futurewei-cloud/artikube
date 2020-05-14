package router

import (
	"fmt"
	"net/http"
	"time"

	arti_logger "github.com/futurewei-cloud/artikube/pkg/server/logger"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
	Action  string
}

type Router struct {
	*gin.Engine
	Logger        *arti_logger.Logger
	Routes        []*Route
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
}

type RouterOption struct {
	Logger      *arti_logger.Logger
	Username    string
	Password    string
	ContextPath string
}

func NewRouter(options RouterOption) *Router {
	engine := gin.New()

	//Add all optional middlewares. The control follow is:
	//Request -> Route Parser -> Middleware -> Route Handler -> middleware - Reponse
	engine.Use(gin.Recovery())

	router := &Router{
		Engine: engine,
		Logger: options.Logger,
		Routes: []*Route{},
	}
	return router
}

func (router *Router) Start(port int) {
	router.Logger.Debug("Starting Artikube", arti_logger.Int("port", port))

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		ReadTimeout:  router.ReadTimeout,
		WriteTimeout: router.WriteTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		router.Logger.Error(err.Error())
	}
	
}
