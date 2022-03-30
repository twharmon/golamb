package golamb

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
)

func TestServiceProviderDBWithConfig(t *testing.T) {
	sp := &awsServiceProvider{
		config: &AWSServiceProviderConfig{
			DynamoDB: &aws.Config{
				Region: aws.String("bar"),
			},
		},
	}
	db := sp.DynamoDB()
	dynamodb, ok := db.(*dynamodb.DynamoDB)
	if !ok {
		t.Fatalf("expected ok")
	}
	if dynamodb == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderDBNoSPConfig(t *testing.T) {
	sp := &awsServiceProvider{config: &AWSServiceProviderConfig{}}
	db := sp.DynamoDB()
	dynamodb, ok := db.(*dynamodb.DynamoDB)
	if !ok {
		t.Fatalf("expected ok")
	}
	if dynamodb == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSESNoSPConfig(t *testing.T) {
	sp := &awsServiceProvider{config: &AWSServiceProviderConfig{}}
	svc := sp.SES()
	ses, ok := svc.(*ses.SES)
	if !ok {
		t.Fatalf("expected ok")
	}
	if ses == nil {
		t.Fatalf("expected not nil")
	}
}
