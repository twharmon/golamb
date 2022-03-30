package golamb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

type awsServiceProvider struct {
	config   *AWSServiceProviderConfig
	session  *session.Session
	dynamodb *dynamodb.DynamoDB
	ses      *ses.SES
	s3       *s3.S3
}

type AWSServiceProviderConfig struct {
	Default  *aws.Config
	DynamoDB *aws.Config
	SES      *aws.Config
	S3       *aws.Config
}

type AWSServiceProvider interface {
	DynamoDB() dynamodbiface.DynamoDBAPI
	SES() sesiface.SESAPI
	S3() s3iface.S3API
}

func (sp *awsServiceProvider) loadSession() {
	if sp.session != nil {
		return
	}
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}
	sp.session = sess
}

func (sp *awsServiceProvider) loadDynamoDB() {
	if sp.dynamodb != nil {
		return
	}
	sp.loadSession()
	if sp.config.DynamoDB != nil {
		sp.dynamodb = dynamodb.New(sp.session, sp.config.DynamoDB)
		return
	}
	if sp.config.Default != nil {
		sp.dynamodb = dynamodb.New(sp.session, sp.config.Default)
		return
	}
	sp.dynamodb = dynamodb.New(sp.session)
}

func (sp *awsServiceProvider) loadSES() {
	if sp.ses != nil {
		return
	}
	sp.loadSession()
	if sp.config.SES != nil {
		sp.ses = ses.New(sp.session, sp.config.SES)
		return
	}
	if sp.config.Default != nil {
		sp.ses = ses.New(sp.session, sp.config.Default)
		return
	}
	sp.ses = ses.New(sp.session)
}

func (sp *awsServiceProvider) loadS3() {
	if sp.s3 != nil {
		return
	}
	sp.loadSession()
	if sp.config.S3 != nil {
		sp.s3 = s3.New(sp.session, sp.config.S3)
		return
	}
	if sp.config.Default != nil {
		sp.s3 = s3.New(sp.session, sp.config.Default)
		return
	}
	sp.s3 = s3.New(sp.session)
}

func (sp *awsServiceProvider) DynamoDB() dynamodbiface.DynamoDBAPI {
	sp.loadDynamoDB()
	return sp.dynamodb
}

func (sp *awsServiceProvider) SES() sesiface.SESAPI {
	sp.loadSES()
	return sp.ses
}

func (sp *awsServiceProvider) S3() s3iface.S3API {
	sp.loadS3()
	return sp.s3
}
