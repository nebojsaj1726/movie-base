package router

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/nebojsaj1726/movie-base/internal/api/gql"
	"github.com/rs/cors"
)

func NewRouter(resolver *gql.Resolver) http.Handler {
	r := mux.NewRouter()

	gqlHandler := handler.New(gql.NewExecutableSchema(gql.Config{
		Resolvers:  resolver,
		Directives: gql.DirectiveRoot{},
		Complexity: gql.ComplexityRoot{},
	}))

	gqlHandler.AddTransport(transport.POST{})
	gqlHandler.Use(extension.Introspection{})

	r.Handle("/playground", playground.Handler("GraphQL Playground", "/query"))
	r.Handle("/query", gqlHandler)

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("ui/dist/assets"))))
	r.Path("/favicon.png").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/dist/favicon.png")
	}))

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ui/dist/index.html")
	})

	handlerWithCors := cors.AllowAll().Handler(r)

	return handlerWithCors
}
