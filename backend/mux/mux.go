package mux

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/koralle/lazy-warehouse/backend/config"
	"github.com/koralle/lazy-warehouse/backend/graph"
	"github.com/koralle/lazy-warehouse/backend/graph/service"
	"github.com/koralle/lazy-warehouse/backend/internal"
	"github.com/koralle/lazy-warehouse/backend/store"
)

func NewMux(ctx context.Context, config *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	db, cleanup, err := store.New(ctx, config)
	if err != nil {
		return nil, cleanup, err
	}

	service := service.New(db)

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{
		Srv: service,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})

	mux.Post("/query", srv.ServeHTTP)

	if config.Env == "dev" {
		srv.Use(extension.Introspection{})
		mux.Get("/", playground.Handler("GraphQL playground", "/query"))
	}

	return mux, func() { fmt.Println("Clean up.") }, nil

}
