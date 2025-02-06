package database

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	"github.com/dhruvgupta7733/notification-service/model"
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

func DbConnect() (*sql.DB){
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Couldn't load env variables!")
	}
	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/notification",  os.Getenv("DB_PASS"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
	return db
}

func CreateEmailRequest(id string, DB *sql.DB ,in *model.EmailRequest){
	insertQuery := `
	INSERT INTO EmailRequest (id, emailId, message)
	VALUES (?, ?, ?);`

	_, err := DB.Exec(insertQuery, id, in.EmailId, in.Message)
	if err != nil {
		log.Printf("Failed to insert record: %v", err)
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