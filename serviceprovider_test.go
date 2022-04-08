package golamb

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/sts"
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

func TestServiceProviderSSMWithConfig(t *testing.T) {
	sp := &awsServiceProvider{
		config: &AWSServiceProviderConfig{
			SSM: &aws.Config{
				Region: aws.String("bar"),
			},
		},
	}
	svc := sp.SSM()
	ssm, ok := svc.(*ssm.SSM)
	if !ok {
		t.Fatalf("expected ok")
	}
	if ssm == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSSMDefaultSPConfig(t *testing.T) {
	sp := &awsServiceProvider{config: &AWSServiceProviderConfig{
		Default: &aws.Config{
			Region: aws.String("bar"),
		},
	}}
	svc := sp.SSM()
	ssm, ok := svc.(*ssm.SSM)
	if !ok {
		t.Fatalf("expected ok")
	}
	if ssm == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSSMNoSPConfig(t *testing.T) {
	c := &handlerContext{
		sp: &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	svc := c.AWS().SSM()
	ssm, ok := svc.(*ssm.SSM)
	if !ok {
		t.Fatalf("expected ok")
	}
	if ssm == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSSMCached(t *testing.T) {
	c := &handlerContext{
		sp: &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	svc := c.AWS().SSM()
	svc = c.AWS().SSM()
	ssm, ok := svc.(*ssm.SSM)
	if !ok {
		t.Fatalf("expected ok")
	}
	if ssm == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSTSWithConfig(t *testing.T) {
	sp := &awsServiceProvider{
		config: &AWSServiceProviderConfig{
			STS: &aws.Config{
				Region: aws.String("bar"),
			},
		},
	}
	svc := sp.STS()
	sts, ok := svc.(*sts.STS)
	if !ok {
		t.Fatalf("expected ok")
	}
	if sts == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSTSDefaultSPConfig(t *testing.T) {
	sp := &awsServiceProvider{config: &AWSServiceProviderConfig{
		Default: &aws.Config{
			Region: aws.String("bar"),
		},
	}}
	svc := sp.STS()
	sts, ok := svc.(*sts.STS)
	if !ok {
		t.Fatalf("expected ok")
	}
	if sts == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSTSNoSPConfig(t *testing.T) {
	c := &handlerContext{
		sp: &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	svc := c.AWS().STS()
	sts, ok := svc.(*sts.STS)
	if !ok {
		t.Fatalf("expected ok")
	}
	if sts == nil {
		t.Fatalf("expected not nil")
	}
}

func TestServiceProviderSTSCached(t *testing.T) {
	c := &handlerContext{
		sp: &awsServiceProvider{config: &AWSServiceProviderConfig{}},
	}
	svc := c.AWS().STS()
	svc = c.AWS().STS()
	sts, ok := svc.(*sts.STS)
	if !ok {
		t.Fatalf("expected ok")
	}
	if sts == nil {
		t.Fatalf("expected not nil")
	}
}
