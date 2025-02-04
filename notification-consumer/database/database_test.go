package database_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/dhruvgupta7733/notification-consumer/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

func setup() {
	db = database.DbConnect()
	_, err := db.Exec("CREATE TEMPORARY TABLE IF NOT EXISTS EmailRequest ( id BIGINT PRIMARY KEY AUTO_INCREMENT, emailId VARCHAR(255), message TEXT, failure_code BIGINT, failure_comments TEXT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP );")
	if err != nil {
		fmt.Printf("Failed to create test table: %v\n", err)
	}
}

func teardown() {
	if db != nil {
		db.Close()
	}
}

func TestDbConnect(t *testing.T) {
	setup()
	assert.NotNil(t, db, "Database connection should not be nil")
	teardown()
}

func TestUpdateEmailRequest(t *testing.T) {
	setup()
	defer teardown()
	insertQuery := `
	INSERT INTO EmailRequest (id, emailId, message, failure_code, failure_comments)
	VALUES (?, ?, ?, ?, ?);`

	res, err := db.Exec(insertQuery, 123, "sample@example.com", "Sample Message", 123, "no")
	
	assert.NoError(t, err)
	id, _ := res.LastInsertId()

	req := &database.Req{
		Id:               id,
		Failure_code:     1,
		Failure_comments: "Test Failure",
	}

	database.UpdateEmailRequest(db, req)

	updatedReq, err := database.GetEmailRequest(db, id)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), updatedReq.Failure_code)
	assert.Equal(t, "Test Failure", updatedReq.Failure_comments)
}

func TestGetEmailRequest(t *testing.T) {
	setup()
	defer teardown()
	insertQuery := `
	INSERT INTO EmailRequest (id, emailId, message, failure_code, failure_comments)
	VALUES (?, ?, ?, ?, ?);`

	res, err := db.Exec(insertQuery, 123, "sample@example.com", "Sample Message", 123, "no")
	assert.NoError(t, err)
	id, _ := res.LastInsertId()
	req, err := database.GetEmailRequest(db, id)
	assert.NoError(t, err)
	assert.Equal(t, id, req.Id)
	assert.Equal(t, "sample@example.com", req.EmailId)
	assert.Equal(t, "Sample Message", req.Message)
	assert.Equal(t, int64(123), req.Failure_code)
	assert.Equal(t, "no", req.Failure_comments)
}
