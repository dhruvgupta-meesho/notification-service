package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/dhruvgupta7733/notification-consumer/database"
	"github.com/joho/godotenv"
)

func (sc *ServiceContainer) SendMail (msg string, id int64) {

	fmt.Println("string that was received => " + msg)
	godotenv.Load(".env")
	from := "dhruvipul1234@gmail.com"
	password := os.Getenv("EMAIL_PASS")

	toEmailAddress := "dhruvgupta3377@gmail.com"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	sc.checkEmailAddr(address, toEmailAddress, id)

	if !sc.CheckIsBlocked(id){
		subject := "Subject: Notification\n"
		contentType := "Content-Type: text/plain; charset=\"utf-8\"\n"
		incomingReq, _ := database.GetEmailRequest(sc.Db,int64(id))

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
		fmt.Println("Email Sent")
	}else{
		database.UpdateEmailRequest(sc.Db, &database.Req{
			Id : id,
			Failure_code: 103,
			Failure_comments: "Blacklisted Email",
		})
		createIndex(sc.Es, string("Couldn't Send Email(Blacklisted) : " + toEmailAddress))
		fmt.Println("Can't Send Email (Blacklisted)")
	}

}

func (sc *ServiceContainer) checkEmailAddr(address string, toEmailAddress string, id int64){
	c, _ := smtp.Dial(address)

	err := c.Verify(toEmailAddress)
	
	if err!= nil{
		log.Print("Wrong Email Address : ", err)
		createIndex(sc.Es, string("Wrong Email Address : "+toEmailAddress))
		database.UpdateEmailRequest(sc.Db, &database.Req{
			Id : id,
			Failure_code: 102,
			Failure_comments: "Wrong Email Address",
		})
		return
	}
}
