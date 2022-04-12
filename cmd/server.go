package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph"
	"github.com/memeoAmazonas/test-nextjs-golang-graphql-back-2022/graph/generated"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const defaultPort = "3201"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error loading")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	/*	cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			Debug:            true,
		})
		router.Use(cors.Handler)*/

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	handler := cors.Default().Handler(router)
	log.Printf("server graphql connect to http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
