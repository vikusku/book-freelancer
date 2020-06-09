package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

var s *session.Session

func GetAWSSession() *session.Session {
	if s == nil {
		var err error
		s, err = session.NewSession(&aws.Config{
			Region: aws.String("eu-west-1")},
		)

		if err != nil {
			panic(err)
		}

		log.Println("Created aws session")
	}

	return s
}

