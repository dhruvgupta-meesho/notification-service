package services

import (
	"context"
	"fmt"
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

func (sc *ServiceContainer) CheckIsBlocked (id int64) bool {
	time.Sleep(4 * time.Second) 
	ctx := context.Background()
	sc.Rdb.SAdd(ctx, "blocked", "hell")
	v := sc.Rdb.SIsMember(ctx, "blocked", id)
	if v.Val(){
		fmt.Println("It is Present!")	
	}else{
		fmt.Println("It is not Present!")
	}
	return v.Val()
}