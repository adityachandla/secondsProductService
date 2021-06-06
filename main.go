package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

const tablename = "Trial"

type Column struct {
	Key   string `dynamodbav:"key"`
	Value string `dynamodbav:"value"`
}

func main() {
	client := getClient()
	for idx := 0; idx < 100; idx++ {
		myValue := &Column{
			Key:   uuid.New().String(),
			Value: fmt.Sprintf("value -> %d", rand.Int31()),
		}
		addToTable(client, myValue)
	}
}

func getClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		panic(err)
	}
	return dynamodb.NewFromConfig(cfg)
}

func addToTable(client *dynamodb.Client, item *Column) {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		panic(err)
	}

	putItemInput := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tablename),
	}

	_, err = client.PutItem(context.TODO(), putItemInput)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted item")
}
