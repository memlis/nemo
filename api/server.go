package api

import (
 	"github.com/gorilla/mux"
 	"net"
 	"net/http"
	"log"
)

type Server struct {
	addr    string
	routers []router.Router
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func NewServer(addr string) *Server {
	s := &Server {
            addr: addr,
	}

	s.setupRouters()

	return s
}

func (s *Server) setupRouters() {
	s.routers = map[string]map[string]handlerFunc{
                "GET": map[string]handleFunc{
                        "/v1/jobs":          getJobs,
                        "/v1/jobs/:id":          getJob,
                },  
                "POST": map[string]handleFunc{
                        "/jobs":    startJob,
                        "/jobs/:id":      updateJob,
                        "/jobs/:id/scale":      scaleJob,
                },  
                "DELETE": map[string]handleFunc{
                        "/jobs/:id":  delJob,
                },  
        }   
}
func (s *Server) createHandler() *mux.Router {
	m := mux.NewRouter()

	for method, router := range s.routers {
		for path, handler := range router {
			m.Path(path).Methods(method).Handler(handler)
		}
	}

	return m
}

func (s *Server) ListenAndServe() error {
 	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.createHandler(),
	}
	log.Println("API Server listen on %s", s.addr)
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Println("Listen on %s error: %s", s.addr, err)
		return err
	}
	return srv.Serve(ln)
}

