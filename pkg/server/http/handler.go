package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seblegall/sigstat/pkg/sigstat"
)

//Handler actually handle http requests.
//It use a router to map uri to HandlerFunc
type Handler struct {
	client sigstat.Client
	router *mux.Router
}

//Route is an api route
type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is a list of web server routes
type routes []route

//NewHandler create an Handler using defined routes.
//It takes a client as argument in order to be passe to the handler and be accessible to the HandlerFunc
//Typically in a CRUD API, the client manage connections to a storage system.
func NewHandler(c sigstat.Client) *Handler {
	h := &Handler{
		client: c,
	}

	var routes = routes{
		route{
			"UpdateStatus",
			"PATCH",
			"/status/",
			h.UpdateStatus,
		},
	}

	h.router = newRouter(routes)

	return h
}

//newRouter create a new router for the webserver, used by the handler
func newRouter(rtes routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, r := range rtes {
		router.
			Methods(r.Method).
			Path(r.Pattern).
			Name(r.Name).
			Handler(r.HandlerFunc)
	}

	return router
}

//Router returns the defined router for the Handler
func (h *Handler) Router() *mux.Router { return h.router }

//UpdateStatus handler PUT request that update the status for a given Command ID.
func (h *Handler) UpdateStatus(w http.ResponseWriter, r *http.Request) {

	cmd := sigstat.Command{
		Status: "running",
	}

	h.client.CommandService().UpdateStatus(cmd)
}
