package main

import (
	"log"
	"os"

	"github.com/kytruong0712/go-market/product-service/api/cmd/banner"
	"github.com/kytruong0712/go-market/product-service/api/internal/config/db/pg"
)

func main() {
	banner.Print()

	// Initial DB connection
	conn, err := pg.Connect(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}
