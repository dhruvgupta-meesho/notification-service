package services

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/hashicorp/go-uuid"
	"github.com/joho/godotenv"
)

func getESClient() *elasticsearch.Client{
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Couldn't load env variables!")
	}
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:" + os.Getenv("ELASTIC_PORT"),
		},
		Password: os.Getenv("ELASTIC_PASS"),
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil{
		log.Printf("Unable to connect to Elasticsearch!")
	}
	log.Printf("Connected to Elasticsearch!")
	return es
}

func createIndex(es *elasticsearch.Client, data string){
	uid, _ := uuid.GenerateUUID()
	request:= esapi.IndexRequest{Index: "Noify", DocumentID: uid, Body: strings.NewReader(data)}
	request.Do(context.Background(), es)

	log.Printf("Created elasticsearch Index!")
}
