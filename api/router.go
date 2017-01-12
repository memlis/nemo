package api

import (
	"net/http"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

type Routes map[string]map[string]map[string]handlerFunc

type Router struct {
	routes Routes
	store  Store
}

func NewRouter(s Store) *Router {
	r := &Router{
		store: s,
	}

	r.initRoutes()
	return r
}

func (r *Router) initRoutes() {
	r.routes = Routes{
		"v1": {
			"GET": {
				"/job/all":     getAllJobs,
				"/job/{id:.+}": getJob,
			},
			"POST": {
				"/job/new": startJob,
			},
			"DELETE": {
				"/job/all":     delAllJobs,
				"/job/{id:.+}": delJob,
			},
		},
	}
}
