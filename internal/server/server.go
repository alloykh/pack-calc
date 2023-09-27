package server

import (
	"github.com/alloykh/pack-calc/internal/query"
	"github.com/alloykh/pack-calc/pkg/log"
	"github.com/alloykh/pack-calc/pkg/server"
	"github.com/gorilla/mux"
)

const (
	base = "/api/v1/"

	calculatePacksEndpoint = base + "calc/packs"
)

func NewHttpHandler(logger log.Logger, resolver *query.Resolver) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc(calculatePacksEndpoint, server.HandleJSONPost(resolver.CalcPacks))

	return r
}
