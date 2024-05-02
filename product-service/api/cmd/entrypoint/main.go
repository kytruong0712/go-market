package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/kytruong0712/go-market/product-service/api/cmd/banner"
	"github.com/kytruong0712/go-market/product-service/api/internal/config/db/pg"
	"github.com/kytruong0712/go-market/product-service/api/internal/config/httpserver"
	"github.com/kytruong0712/go-market/product-service/api/internal/controller/categories"
	"github.com/kytruong0712/go-market/product-service/api/internal/controller/system"
	systemrest "github.com/kytruong0712/go-market/product-service/api/internal/handler/rest/system"
	"github.com/kytruong0712/go-market/product-service/api/internal/repository"
)

func main() {
	banner.Print()

	ctx := context.Background()

	// Initial DB connection
	conn, err := pg.Connect(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Initial router
	rtr, err := initRouter(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	httpserver.Start(httpserver.Handler(
		ctx,
		httpserver.NewCORSConfig(rtr.corsOrigins),
		systemrest.New(rtr.systemCtrl).CheckDBReady(),
		rtr.routes))
}

func initRouter(
	ctx context.Context,
	db *sql.DB) (router, error) {
	repo := repository.New(db)

	systemCtrl := system.New(repo)
	categoryCtrl := categories.New(repo)

	return router{
		systemCtrl:   systemCtrl,
		categoryCtrl: categoryCtrl,
		corsOrigins:  []string{"*"},
	}, nil
}
