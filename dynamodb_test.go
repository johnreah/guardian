package guardian

import (
	"fmt"
	"time"
	"testing"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

func TestListTables(t *testing.T) {
	fmt.Println("Starting...")
	startTime := time.Now()

	ListTables()

	fmt.Printf("\nFinished in %dms\n", time.Now().Sub(startTime)/1000000)
}

func TestCreateTable(t *testing.T) {
	input := dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("year"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("year"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("title"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Movies"),
	}
	err := CreateTable(&input)
	if err != nil {
		t.Errorf("Failed to create table: %v", err)
	}
}

func TestDeleteTable(t *testing.T) {
	input := dynamodb.DeleteTableInput{
		TableName: aws.String("Movies"),
	}
	err := DeleteTable(&input)
	if err != nil {
		t.Errorf("Failed to delete table: %v", err)
	}
}