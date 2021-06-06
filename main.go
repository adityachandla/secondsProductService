package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		panic(err)
	}
	dynamoDbClient := dynamodb.NewFromConfig(cfg)
	resp, err := dynamoDbClient.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		panic(err)
	}
	for _, name := range resp.TableNames {
		fmt.Printf("Name: %s\n", name)
	}
}
