package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/nodes", queryNodesHandlerFunc(ctx)).
		Methods("GET")
	router.HandleFunc("/nodes/{address}", queryNodeHandlerFunc(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/nodes", txRegisterNodeHandlerFunc(ctx)).
		Methods("POST")
}

func RegisterRoutes(ctx context.CLIContext, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}
