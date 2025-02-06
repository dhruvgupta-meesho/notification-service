package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Req struct {
	Id               string
	EmailId          string
	Message          string
	Failure_code     int64
	Failure_comments string
	Created_at       string
	Updated_at       string
}

func DbConnect() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't load env variables! ", err)
	}
	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/notification",  os.Getenv("DB_PASS"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func UpdateEmailRequest(db *sql.DB, r *Req)  {
	_, err := db.Exec("UPDATE EmailRequest SET failure_code = ?, failure_comments = ? WHERE id = ?", r.Failure_code, r.Failure_comments, r.Id)
	if err != nil {
		log.Printf("Failed to execute update query: %v", err)
	}
}

func GetEmailRequest(db *sql.DB, id string)(Req, error){
	var r Req
	err := db.QueryRow("SELECT * FROM EmailRequest WHERE id = ?", id).Scan(&r.Id, &r.EmailId, &r.Message, &r.Failure_code, &r.Failure_comments, &r.Created_at, &r.Updated_at)
	if err != nil {
		log.Printf("Error : %s", err)
		return r, err
	}
	fmt.Println(r)
	return r, nil
}
