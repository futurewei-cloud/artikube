package artikube

import (
	arti_router "github.com/futurewei-cloud/artikube/pkg/server/router"
)

func (svr *Server) Routes() []*arti_router.Route {
	var routes []*arti_router.Route

	root := arti_router.Route{"GET", "/", svr.getRootPageHandler, ""}
	routes = append(routes, &root)

	return routes
}
