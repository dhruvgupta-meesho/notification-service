package services

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func MakeRedisConn() *redis.Client{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load env variables! ", err)
	}
	
	rdb := redis.NewClient(&redis.Options{
        Addr:	  "localhost:" + os.Getenv("REDIS_PORT"),
        Password: os.Getenv("REDIS_PASS"),
        DB:		  0,
    })
	return rdb
}

func (sc *ServiceContainer) CheckIsBlocked (to string) bool {
	time.Sleep(4 * time.Second) 
	ctx := context.Background()
	v := sc.Rdb.SIsMember(ctx, "blocked", to)
	if v.Val(){
		log.Println("It is Present! ", to)	
	}else{
		log.Println("It is not Present!", to)
	}
	return v.Val()
}