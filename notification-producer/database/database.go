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

func DbConnect() (*sql.DB){
	godotenv.Load(".env")
	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/notification",  os.Getenv("DB_PASS"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
	return db
}

func CreateEmailRequest(DB *sql.DB ,in *model.EmailRequest){
	insertQuery := `
	INSERT INTO EmailRequest (id, emailId, message, failure_code, failure_comments)
	VALUES (?, ?, ?, ?, ?);`

	_, err := DB.Exec(insertQuery, in.Id, in.EmailId, in.Message, in.FailureCode, in.FailureComments)
	if err != nil {
		log.Printf("Failed to insert record: %v", err)
	}
}