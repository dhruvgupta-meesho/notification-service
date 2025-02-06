package database_test

// import (
// 	"database/sql"
// 	"log"
// 	"testing"

// 	"github.com/dhruvgupta7733/notification-service/database"
// 	"github.com/dhruvgupta7733/notification-service/model"
// 	"github.com/stretchr/testify/assert"
// )


// var db *sql.DB

// func setup(t *testing.T) *sql.DB {
// 	db = database.DbConnect()
// 	_, err := db.Exec(`CREATE TEMPORARY TABLE IF NOT EXISTS EmailRequest (
// 		id BIGINT PRIMARY KEY AUTO_INCREMENT,
// 		emailId VARCHAR(255),
// 		message TEXT,
// 		failure_code BIGINT,
// 		failure_comments TEXT,
// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
// 	);`)
// 	assert.NoError(t, err)
// 	t.Cleanup(func() { db.Close() })
// 	return db
// }

// func TestDbConnect(t *testing.T) {
// 	db = setup(t)
// 	assert.NotNil(t, db, "Database connection should not be nil")
// }

// func TestCreateEmailRequest(t *testing.T) {
// 	db = setup(t)

// 	emailRequest := &model.EmailRequest{
// 		Id:              101,
// 		EmailId:         "test@example.com",
// 		Message:         "Test Message",
// 		FailureCode:     0,
// 		FailureComments: "No failure",
// 	}

// 	database.CreateEmailRequest(db, emailRequest)

// 	var count int
// 	err := db.QueryRow(`SELECT COUNT(*) FROM EmailRequest WHERE id = ?`, emailRequest.Id).Scan(&count)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, count, "Record should be inserted successfully")

// 	// Optional: Clean up inserted test data
// 	_, err = db.Exec(`DELETE FROM EmailRequest WHERE id = ?`, emailRequest.Id)
// 	if err != nil {
// 		log.Printf("Failed to clean up test record: %v", err)
// 	}
// }
