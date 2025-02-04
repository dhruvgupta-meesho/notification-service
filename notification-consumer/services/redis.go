package services

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func MakeRedisConn() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "",
        DB:		  0,
    })
	return rdb
}

func (sc *ServiceContainer) CheckIsBlocked (to string) bool {
	time.Sleep(4 * time.Second) 
	ctx := context.Background()
	sc.Rdb.SAdd(ctx, "blocked", "hell")
	v := sc.Rdb.SIsMember(ctx, "blocked", to)
	if v.Val(){
		log.Println("It is Present!")	
	}else{
		log.Println("It is not Present!")
	}
	return v.Val()
}