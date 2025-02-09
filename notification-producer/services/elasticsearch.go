package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/dhruvgupta7733/notification-service/model"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/joho/godotenv"
)

type Eslogs struct{
	Id string `json:"id"`
	Email string `json:"email"`
	Comment string `json:"comment"`
	Timestamp string `json:"timestamp"`
}

func GetESClient() *elasticsearch.Client{
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

func FetchDocuments(es *elasticsearch.Client, logreq *model.LogRequest) [][]byte {

	query := fmt.Sprintf(`{
		"query": {
			"bool": {
				"must": [
					{
						"term": {
							"email.keyword": "%s"
						}
					},
					{
						"range": {
							"timestamp": {
								"time_zone": "+01:00",
								"gte": "%s",
								"lte": "%s"
							}
						}
					}
				]
			}
		}
	}`, logreq.Email, logreq.Start, logreq.End)
	

	req := esapi.SearchRequest{
		Index: []string{"notify"},
		Body:  bytes.NewReader([]byte(query)),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error fetching documents: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response from Elasticsearch: %s", res.String())
	}

	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Error parsing response: %s", err)
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	resp := [][]byte{}

	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		data, _ := json.MarshalIndent(source, "", "  ")
		resp = append(resp, data)
	}
	return resp
}
