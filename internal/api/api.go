package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"hw36a.4.1/internal/postgres"
)

type API struct {
	db     *postgres.Store
	router *mux.Router
}

func New(db postgres.Store) *API {
	api := API{
		db: &db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) endpoints() {
	api.router.HandleFunc("/news/{n}", api.postsHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./cmd/server/webapp"))))
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) postsHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["n"]
	n, err := strconv.Atoi(s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	posts, err := api.db.Posts(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(bytes)
}
