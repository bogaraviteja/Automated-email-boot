package main

import (
	"encoding/csv"
	"log"
	"net/smtp"
	"os"
)

type Person struct {
	Name  string
	Email string
}

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func main() {
	importPersons("email.csv")
	sendEmail()
}

func importPersons(file string) []Person {
	var ps []Person

	f, err := os.Open(file)
	if err != nil {
		log.Fatalln("Could not open file", err)
	}

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("Could not read CSV", err)
	}

	for i, v := range rows {
		if i == 0 {
			continue
		}

		ps = append(ps, Person{
			Name:  v[0],
			Email: v[1],
		})
	}

	return ps
}

func getMessageString(from, to, subject, body, signature string) []byte {
	return []byte("From: " + from + "\r\n" + "To: " + to + "\r\n" + "Subject: " + subject + "\r\n" + body + "\r\n" + signature + "\r\n")
}

func sendEmail() {
	ps := importPersons("email.csv")

	for _, v := range ps {
		from := "ravitejaboga336@gmail.com"
		appPass := "*****sender passwors***"
		to := []string{
			v.Email,
		}
		s := smtpServer{host: "smtp.gmail.com", port: "587"}
		b := getMessageString(from, to[0], "Hello", "Hello, this is test mail from automated code.", "Thank you!")
		auth := smtp.PlainAuth("", from, appPass, s.host)

		err := smtp.SendMail(s.Address(), auth, from, to, b)
		if err != nil {
			log.Fatalln("Could not send email", err)
		}
	}
}
