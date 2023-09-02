package main

import (
	"log"
	"testing-nextalent/config"
	"testing-nextalent/database"
	"testing-nextalent/external"
	"testing-nextalent/pkg/httpclient"

	templateData "testing-nextalent/data"
	templateServer "testing-nextalent/handler"
	templateHandler "testing-nextalent/handler/api"
	templateService "testing-nextalent/service"
)

func main() {

	err := config.Init("config.yaml")
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}

	db := database.Init()

	httpc := httpclient.NewClient()

	ad := external.New(httpc)

	sd := templateData.New(db)
	ss := templateService.New(sd, ad)
	sh := templateHandler.New(ss)

	s := templateServer.Server{
		Api: sh,
	}

	s.Serve(":8080")
}
