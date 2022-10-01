package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type MyVisitorCounter struct {
	Path string // Hash key, a.k.a. partition key

	VisitorCount int
}

func GetAndUpdateCount() (int, error) {

	db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	table := db.Table("MyVisitorCounter")

	var result MyVisitorCounter
	err := table.Update("Path", "countergolang").
		SetExpr("SET VisitorCount = VisitorCount + ?", 1).
		Value(&result)

	return result.VisitorCount, err

}

func HandleRequest(ctx context.Context) (string, error) {
	count, _ := GetAndUpdateCount()
	return fmt.Sprintln(count), nil
}

func main() {
	lambda.Start(HandleRequest)
}
