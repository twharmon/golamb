package golamb

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
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

func TestServiceProviderDBDefaultSPConfig(t *testing.T) {
	sp := &awsServiceProvider{config: &AWSServiceProviderConfig{
		Default: &aws.Config{
			Region: aws.String("bar"),
		},
	}}
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

func TestServiceProviderDBCached(t *testing.T) {
	c := &handlerContext{
		sp: &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	svc := c.AWS().DynamoDB()
	svc = c.AWS().DynamoDB()
	db, ok := svc.(*dynamodb.DynamoDB)
	if !ok {
		t.Fatalf("expected ok")
	}
	if db == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSESWithConfig(t *testing.T) {
	sp := &awsServiceProvider{
		config: &AWSServiceProviderConfig{
			SES: &aws.Config{
				Region: aws.String("bar"),
			},
		},
	}
	svc := sp.SES()
	ses, ok := svc.(*ses.SES)
	if !ok {
		t.Fatalf("expected ok")
	}
	if ses == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSESDefaultSPConfig(t *testing.T) {
	sp := &awsServiceProvider{config: &AWSServiceProviderConfig{
		Default: &aws.Config{
			Region: aws.String("bar"),
		},
	}}
	svc := sp.SES()
	ses, ok := svc.(*ses.SES)
	if !ok {
		t.Fatalf("expected ok")
	}
	if ses == nil {
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

func TestServiceProviderSESCached(t *testing.T) {
	c := &handlerContext{
		sp: &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	svc := c.AWS().SES()
	svc = c.AWS().SES()
	ses, ok := svc.(*ses.SES)
	if !ok {
		t.Fatalf("expected ok")
	}
	if ses == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderS3WithConfig(t *testing.T) {
	sp := &awsServiceProvider{
		config: &AWSServiceProviderConfig{
			S3: &aws.Config{
				Region: aws.String("bar"),
			},
		},
	}
	svc := sp.S3()
	s3, ok := svc.(*s3.S3)
	if !ok {
		t.Fatalf("expected ok")
	}
	if s3 == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderS3DefaultSPConfig(t *testing.T) {
	sp := &awsServiceProvider{config: &AWSServiceProviderConfig{
		Default: &aws.Config{
			Region: aws.String("bar"),
		},
	}}
	svc := sp.S3()
	s3, ok := svc.(*s3.S3)
	if !ok {
		t.Fatalf("expected ok")
	}
	if s3 == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderS3NoSPConfig(t *testing.T) {
	c := &handlerContext{
		sp: &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	svc := c.AWS().S3()
	s3, ok := svc.(*s3.S3)
	if !ok {
		t.Fatalf("expected ok")
	}
	if s3 == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderS3Cached(t *testing.T) {
	c := &handlerContext{
		sp: &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	svc := c.AWS().S3()
	svc = c.AWS().S3()
	s3, ok := svc.(*s3.S3)
	if !ok {
		t.Fatalf("expected ok")
	}
	if s3 == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSessionCached(t *testing.T) {
	c := &handlerContext{
		sp: &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	c.AWS().DynamoDB()
	svc := c.AWS().S3()
	s3, ok := svc.(*s3.S3)
	if !ok {
		t.Fatalf("expected ok")
	}
	if s3 == nil {
		t.Fatalf("expected not nil")
	}
}
