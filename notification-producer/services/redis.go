package services

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func MakeRedisConn() *redis.Client{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load env variables! ", err)
	}

	rdb := redis.NewClient(&redis.Options{
        Addr:	  "localhost:"+ os.Getenv("REDIS_PORT"),
        Password: os.Getenv("REDIS_PASS"),
        DB:		  0,
    })
	return rdb
}


// func AddBlacklistEmail(rdb *redis.Client, email string) string{
// 	ctx := context.Background()
// 	res := rdb.SAdd(ctx, "blocked", email)
// 	if res.Val() == 0{
// 		log.Println("Unsuccessful")
// 		return "Unsuccessful"
// 	}else if res.Val() == 1{
// 		log.Println("Successful")
// 		return "Successful"
// 	}else{
// 		log.Println("Error while blacklisting")
// 		return "Error while blacklisting"
// 	}
// }

// func RemoveBlacklistEmail(rdb *redis.Client, email string) string {
// 	ctx := context.Background()
// 	res := rdb.SRem(ctx, "blocked", email)
// 	log.Println(res)
// 	if res.Val() == 0{
// 		log.Println("Unsuccessful")
// 		return "Unsuccessful"
// 	}else if res.Val() == 1{
// 		log.Println("Successful")
// 		return "Successful"
// 	}else{
// 		log.Println("Error while blacklisting")
// 		return "Error while blacklisting"
// 	}
// }

func AddBlacklistEmails(rdb *redis.Client, emails []string) string {
	ctx := context.Background()
	res := rdb.SAdd(ctx, "blocked", emails)
	if res.Err() != nil {
		log.Println("Error while blacklisting:", res.Err())
		return "Error while blacklisting"
	}
	if res.Val() == 0 {
		log.Println("No new emails added to the blacklist")
		return "No new emails added to the blacklist"
	}
	log.Println("Successfully blacklisted emails")
	GetAllBlacklistedEmails(rdb)
	return "Successfully blacklisted emails"
}

func RemoveBlacklistEmails(rdb *redis.Client, emails []string) string {
	ctx := context.Background()
	res := rdb.SRem(ctx, "blocked", emails)
	if res.Err() != nil {
		log.Println("Error while removing from blacklist:", res.Err())
		return "Error while removing from blacklist"
	}
	if res.Val() == 0 {
		log.Println("No emails found for removal")
		return "No emails found for removal"
	}
	log.Println("Successfully removed emails from blacklist")
	return "Successfully removed emails from blacklist"
}


func GetAllBlacklistedEmails(rdb *redis.Client)[]string{
	ctx := context.Background()
	res, err := rdb.SMembers(ctx, "blocked").Result()
	if err!= nil{
		log.Printf("Error Getting all the blocked ids ", err)
	}
	return res
}