package main

import (
	"etp/pkg/api"
	"etp/pkg/db"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dbase := db.New()
	ps := db.PreSearch{
		EtpID:    "EtpID",
		Article:  "Article",
		Brand:    "Brand",
		PartName: "PartName",
	}
	dbase.NewPreSearch(ps)
	api := api.New(dbase)
	err = http.ListenAndServe(":80", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
