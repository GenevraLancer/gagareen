// Package implements a client for ElastiSearch service
package indexer

import (
	"log"

	config "github.com/GenevraLancer/gagarin/backend/config"
	elastic "github.com/elastic/go-elasticsearch/v7"
	
)

var client *elastic.Client

//Инициализация клиента с параметрами из .env
func init() {
	var err error
	cfg := config.InitConfig()
	client, err = elastic.NewClient(*cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %v", err)
	}

	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}
