package main

import (
	"log"

	"github.com/zefrenchwan/m3.git/properties"
	"github.com/zefrenchwan/m3.git/services"
	"github.com/zefrenchwan/m3.git/storage"
)

func main() {
	// load config, stop app if failure
	log.Println("Loading configuration")
	config, loadingErr := properties.LoadLocalProperties("config.properties")
	if loadingErr != nil {
		log.Fatal("Failure to load config for app")
		return
	}

	// Load DAO and display type
	var dao storage.Dao
	if result, err := storage.InitDao(config); err != nil {
		log.Fatal("Cannot start DAO")
		return
	} else {
		dao = result
		log.Println("Started DAO as", dao.Info())
	}

	// DAO is up, start server
	addr := config["SERVING_ADDRESS"]
	log.Println("Starting application")
	log.Println("Will serve to", addr)
	services.LaunchHandler(addr, dao)
}
