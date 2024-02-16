package router

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/nebojsaj1726/movie-base/internal/api/gql"
	"github.com/rs/cors"
)

func NewRouter(resolver *gql.Resolver) http.Handler {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("ui/dist")))
	assetsFs := http.FileServer(http.Dir("ui/dist/assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", assetsFs))

	gqlHandler := handler.New(gql.NewExecutableSchema(gql.Config{
		Resolvers:  resolver,
		Directives: gql.DirectiveRoot{},
		Complexity: gql.ComplexityRoot{},
	}))

	gqlHandler.AddTransport(transport.POST{})

	r.Handle("/playground", playground.Handler("GraphQL Playground", "/query"))
	r.Handle("/query", gqlHandler)

	handlerWithCors := cors.AllowAll().Handler(r)

	return handlerWithCors
}
