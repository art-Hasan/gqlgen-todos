package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/art-Hasan/gqlgen-todos/ent"
	"github.com/art-Hasan/gqlgen-todos/graph/generated"
)

type DB struct {
	Engine string
	URI    string
}

type Addr struct {
	Host string
	Port int
}

func (a Addr) Address() string {
	return net.JoinHostPort(a.Host, strconv.Itoa(a.Port))
}

type service struct {
	ent    *ent.Client
	router *chi.Mux

	addr Addr
}

func (s service) ListenAndServe() error {
	log.Printf("listening on %s", s.addr.Address())
	return http.ListenAndServe(s.addr.Address(), s.router)
}

func newService(ent *ent.Client, r *chi.Mux, addr Addr) service {
	return service{ent: ent, router: r, addr: addr}
}

func newAddr() Addr {
	return Addr{
		Host: defaultHost,
		Port: defaultPort,
	}
}

func newDB() DB {
	return DB{
		Engine: "sqlite3",
		URI:    "file:ent?mode=memory&cache=shared&_fk=1",
	}
}

func newEnt(db DB) (*ent.Client, error) {
	client, err := ent.Open(
		db.Engine,
		db.URI,
		ent.Log(log.Print),
	)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed open connection to: %s:%s", db.Engine, db.URI))
	}
	log.Printf("successfully connected to: %s:%s", db.Engine, db.URI)
	return client, nil
}

func newRouter(gr generated.ResolverRoot) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{
		Resolvers: gr,
	})))
	return router
}
