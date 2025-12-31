package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Rahul-aithal/ThumbPicker/db"
	"github.com/Rahul-aithal/ThumbPicker/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Source - https://stackoverflow.com/q/65880069
// Posted by Darien Miller
// Retrieved 2025-12-14, License - CC BY-SA 4.0

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "admin"
		dbname   = "thumbpicker"
	)

	// Create the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	ctx := context.Background()
	conn, err := pgxpool.New(ctx, psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	err = conn.Ping(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	queries := db.New(conn)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", routers.Routers(queries))
	filesDir := http.Dir(filepath.Base("./pub"))
	fileServer(r, "/pub", filesDir)
	fmt.Println("Server is running at :4000")
	http.ListenAndServe(":4000", r)
}

func fileServer(r chi.Router, serverRoute string, pathToStaticFolder http.FileSystem) {
	if strings.ContainsAny(serverRoute, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if serverRoute != "/" && serverRoute[len(serverRoute)-1] != '/' {
		r.Get(serverRoute, http.RedirectHandler(serverRoute+"/", 301).ServeHTTP)
		serverRoute += "/"
	}
	serverRoute += "*"

	r.Get(serverRoute, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		serverRoutePrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(serverRoutePrefix, http.FileServer(pathToStaticFolder))
		fs.ServeHTTP(w, r)
	})
}
