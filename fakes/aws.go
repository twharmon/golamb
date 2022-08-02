package fakes

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/aws/aws-sdk-go/service/sfn/sfniface"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
)

// AWS implements the golamb.AWSServiceProvider interface.
type AWS struct {
	dynamodb dynamodbiface.DynamoDBAPI
	ses      sesiface.SESAPI
	s3       s3iface.S3API
	sts      stsiface.STSAPI
	ssm      ssmiface.SSMAPI
	sfn      sfniface.SFNAPI
}

// NewAWS creates a value that implements the
// golamb.AWSServiceProvider interface.
func NewAWS() *AWS {
	return &AWS{}
}

// DynamoDB implements the golamb.AWSServiceProvider interface.
func (a *AWS) DynamoDB() dynamodbiface.DynamoDBAPI {
	return a.dynamodb
}

// SES implements the golamb.AWSServiceProvider interface.
func (a *AWS) SES() sesiface.SESAPI {
	return a.ses
}

// S3 implements the golamb.AWSServiceProvider interface.
func (a *AWS) S3() s3iface.S3API {
	return a.s3
}

// STS implements the golamb.AWSServiceProvider interface.
func (a *AWS) STS() stsiface.STSAPI {
	return a.sts
}

// SSM implements the golamb.AWSServiceProvider interface.
func (a *AWS) SSM() ssmiface.SSMAPI {
	return a.ssm
}

// SFN implements the golamb.AWSServiceProvider interface.
func (a *AWS) SFN() sfniface.SFNAPI {
	return a.sfn
}

// WithDynamoDB sets the dynamodb client of the AWSServiceProvider.
func (a *AWS) WithDynamoDB(svc dynamodbiface.DynamoDBAPI) *AWS {
	a.dynamodb = svc
	return a
}

// WithSES sets the ses client of the AWSServiceProvider.
func (a *AWS) WithSES(svc sesiface.SESAPI) *AWS {
	a.ses = svc
	return a
}

// WithS3 sets the s3 client of the AWSServiceProvider.
func (a *AWS) WithS3(svc s3iface.S3API) *AWS {
	a.s3 = svc
	return a
}

// WithSTS sets the sts client of the AWSServiceProvider.
func (a *AWS) WithSTS(svc stsiface.STSAPI) *AWS {
	a.sts = svc
	return a
}

// WithSSM sets the ssm client of the AWSServiceProvider.
func (a *AWS) WithSSM(svc ssmiface.SSMAPI) *AWS {
	a.ssm = svc
	return a
}

// WithSFN sets the sfn client of the AWSServiceProvider.
func (a *AWS) WithSFN(svc sfniface.SFNAPI) *AWS {
	a.sfn = svc
	return a
}
