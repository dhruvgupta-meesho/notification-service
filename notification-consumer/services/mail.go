package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"

	"github.com/dhruvgupta7733/notification-consumer/database"
	"github.com/joho/godotenv"
)

func (sc *ServiceContainer) SendMail (msg string, id int64) {

	err := godotenv.Load(".env")
	
	if err != nil{
		log.Fatal("Couldn't load env variables!")
	}

	from := "dhruvipul1234@gmail.com"
	password := os.Getenv("EMAIL_PASS")

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	incomingReq, _ := database.GetEmailRequest(sc.Db,int64(id))
	
	// toEmailAddress := []string{incomingReq.EmailId}
	toEmailAddress := "dhruvgupta3377@gmail.com"
	to := []string{toEmailAddress}

	sc.checkEmailAddr(address, toEmailAddress, id)

	if !sc.CheckIsBlocked(incomingReq.EmailId){
		subject := "Subject: Notification\n"
		contentType := "Content-Type: text/plain; charset=\"utf-8\"\n"

		body := incomingReq.Message
		message := []byte(subject + contentType + "\n" + body)
		
		auth := smtp.PlainAuth("", from, password, host)

		err := smtp.SendMail(address, auth, from, to, message)

		if err != nil {
			database.UpdateEmailRequest(sc.Db, &database.Req{
				Id : id,
				Failure_code: 104,
				Failure_comments: fmt.Sprintf("Internal Server Error(%s)", err),
			})
			log.Print(err)	
		}

		database.UpdateEmailRequest(sc.Db, &database.Req{
			Id : id,
			Failure_code: 101,
			Failure_comments: "Email Sent Successfully",
		})

		createIndex(sc.Es, string("Email Sent Successfully : " + toEmailAddress))
		log.Println("Email Sent!")
	}else{
		database.UpdateEmailRequest(sc.Db, &database.Req{
			Id : id,
			Failure_code: 103,
			Failure_comments: "Blacklisted Email",
		})
		createIndex(sc.Es, string("Couldn't Send Email(Blacklisted) : " + toEmailAddress))
		log.Println("Can't Send Email (Blacklisted)")
	}

}

// func (sc *ServiceContainer) checkEmailAddr(address string, toEmailAddress string, id int64){
// 	c, _ := smtp.Dial(address)

// 	err := c.Verify(toEmailAddress)
	
// 	if err!= nil{
// 		log.Println("Wrong Email Address : ", err)
// 		createIndex(sc.Es, string("Wrong Email Address : "+ toEmailAddress))
// 		database.UpdateEmailRequest(sc.Db, &database.Req{
// 			Id : id,
// 			Failure_code: 102,
// 			Failure_comments: "Wrong Email Address",
// 		})
// 		return
// 	}
// }


func (sc *ServiceContainer) checkEmailAddr(address string, toEmailAddress string, id int64) {
	host := address
	fromEmail := "dhruvipul1234@gmail.com"
	password := os.Getenv("EMAIL_PASS")

	hostPart, _, _ := net.SplitHostPort(host)

	conn, err := smtp.Dial(host)
	if err != nil {
		log.Println("Failed to connect to SMTP server:", err)
		return
	}
	defer conn.Close()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         hostPart,
	}

	if err := conn.StartTLS(tlsConfig); err != nil {
		log.Println("Failed to initiate STARTTLS:", err)
		return
	}

	auth := smtp.PlainAuth("", fromEmail, password, hostPart)
	if err := conn.Auth(auth); err != nil {
		log.Println("Authentication failed:", err)
		return
	}

	if err := conn.Mail(fromEmail); err != nil {
		log.Println("Failed to set sender:", err)
		return
	}

	if err := conn.Rcpt(toEmailAddress); err != nil {
		log.Println("Invalid recipient email address:", err)
		updateFailure(sc, toEmailAddress, id, "Invalid recipient")
		return
	}

	log.Println("Email address appears valid:", toEmailAddress)
}

func updateFailure(sc *ServiceContainer, toEmailAddress string, id int64, comment string) {
	createIndex(sc.Es, "Error: "+toEmailAddress)
	database.UpdateEmailRequest(sc.Db, &database.Req{
		Id:              id,
		Failure_code:    102,
		Failure_comments: comment,
	})
}

