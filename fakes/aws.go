package fakes

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/aws/aws-sdk-go/service/sfn/sfniface"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

type AWS struct {
	dynamodb dynamodbiface.DynamoDBAPI
	ses      sesiface.SESAPI
	s3       s3iface.S3API
	sts      stsiface.STSAPI
	ssm      ssmiface.SSMAPI
	sfn      sfniface.SFNAPI
}

func NewAWS() *AWS {
	return &AWS{}
}

func (a *AWS) DynamoDB() dynamodbiface.DynamoDBAPI {
	return a.dynamodb
}

func (a *AWS) SES() sesiface.SESAPI {
	return a.ses
}

func (a *AWS) S3() s3iface.S3API {
	return a.s3
}

func (a *AWS) STS() stsiface.STSAPI {
	return a.sts
}

func (a *AWS) SSM() ssmiface.SSMAPI {
	return a.ssm
}

func (a *AWS) SFN() sfniface.SFNAPI {
	return a.sfn
}

func (a *AWS) WithDynamoDB(svc dynamodbiface.DynamoDBAPI) *AWS {
	a.dynamodb = svc
	return a
}

func (a *AWS) WithSES(svc sesiface.SESAPI) *AWS {
	a.ses = svc
	return a
}

func (a *AWS) WithS3(svc s3iface.S3API) *AWS {
	a.s3 = svc
	return a
}

func (a *AWS) WithSTS(svc stsiface.STSAPI) *AWS {
	a.sts = svc
	return a
}

func (a *AWS) WithSSM(svc ssmiface.SSMAPI) *AWS {
	a.ssm = svc
	return a
}

func (a *AWS) WithSFN(svc sfniface.SFNAPI) *AWS {
	a.sfn = svc
	return a
}
