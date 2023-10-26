package main

import (
	"fmt"

	"github.com/TwiN/go-color"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func listIAMUsers() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println("Error Section AWS:", err)
		return
	}

	svc := iam.New(sess)
	input := &iam.ListUsersInput{}
	result, err := svc.ListUsers(input)

	if err != nil {
		fmt.Println("Error List IAM:", err)
		return
	}
	fmt.Println(color.With(color.Red, "Account IAM Users:"))
	for _, user := range result.Users {
		fmt.Println(color.With(color.Blue, *user.UserName))
	}
}

func main() {
	listIAMUsers()
}
