package services

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/hashicorp/go-uuid"
	"github.com/joho/godotenv"
)

type Eslogs struct{
	Id string `json:"id"`
	Email string `json:"email"`
	Comment string `json:"comment"`
	Timestamp string `json:"timestamp"`
}

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
    // var buf bytes.Buffer
	// json.NewEncoder(&buf).Encode()
	eslog := Eslogs{
		Id: uid,
		Comment: data,
		Email: "dhruvgupta3377@gmail.com",
		Timestamp: time.Now().GoString(),
	}
	s, _:= json.Marshal(eslog)
	request:= esapi.IndexRequest{Index: "notify", DocumentID: uid, Body: bytes.NewReader(s)}
	request.Do(context.Background(), es)

	log.Printf("Created elasticsearch Index!")
}
