package guardian

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

var sess = session.Must(session.NewSession())

func reportError(err error) {
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case dynamodb.ErrCodeInternalServerError:
			fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
		default:
			fmt.Println(aerr.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}

func ListTables() (*dynamodb.ListTablesOutput, error){
	svc := dynamodb.New(sess)
	input := &dynamodb.ListTablesInput{}
	result, err := svc.ListTables(input)
	if err != nil {
		reportError(err)
		return nil, err
	}
	fmt.Print(result)
	return result, nil
}

 func CreateTable(input *dynamodb.CreateTableInput) error {
	 svc := dynamodb.New(sess)
	 _, err := svc.CreateTable(input)
	 return err
 }

func DeleteTable(input *dynamodb.DeleteTableInput) error {
	svc := dynamodb.New(sess)
	_, err := svc.DeleteTable(input)
	return err
}

