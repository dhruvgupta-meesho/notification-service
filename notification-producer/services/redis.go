package services

import (
	"context"
	"log"

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


func AddBlacklistEmail(rdb *redis.Client, email string) string{
	ctx := context.Background()
	res := rdb.SAdd(ctx, "blocked", email)
	if res.Val() == 0{
		log.Println("Unsuccessful")
		return "Unsuccessful"
	}else if res.Val() == 1{
		log.Println("Successful")
		return "Successful"
	}else{
		log.Println("Error while blacklisting")
		return "Error while blacklisting"
	}
}

func RemoveBlacklistEmail(rdb *redis.Client, email string) string {
	ctx := context.Background()
	res := rdb.SRem(ctx, "blocked", email)
	log.Println(res)
	if res.Val() == 0{
		log.Println("Unsuccessful")
		return "Unsuccessful"
	}else if res.Val() == 1{
		log.Println("Successful")
		return "Successful"
	}else{
		log.Println("Error while blacklisting")
		return "Error while blacklisting"
	}
}