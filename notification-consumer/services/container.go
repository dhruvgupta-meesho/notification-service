package services

import (
	"database/sql"

	"github.com/dhruvgupta7733/notification-consumer/database"
	"github.com/redis/go-redis/v9"
	"github.com/elastic/go-elasticsearch/v8"

	_ "github.com/go-sql-driver/mysql"
)

type ServiceContainer struct{
	Rdb *redis.Client
	Db  *sql.DB
	Es  *elasticsearch.Client
}

func MakeContainer() *ServiceContainer{
	sc := &ServiceContainer{
		Rdb: MakeRedisConn(),
		Db : database.DbConnect(),
		Es : getESClient(),
	}
	return sc
}