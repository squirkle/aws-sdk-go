package sts_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/sts"
	"github.com/stretchr/testify/assert"
)

var svc = sts.New(&aws.Config{
	Region:      "mock-region",
	Credentials: aws.DetectCreds("AKID", "SECRET", ""),
})

func TestUnsignedRequest_AssumeRoleWithSAML(t *testing.T) {
	req, _ := svc.AssumeRoleWithSAMLRequest(&sts.AssumeRoleWithSAMLInput{
		PrincipalARN:  aws.String("ARN"),
		RoleARN:       aws.String("ARN"),
		SAMLAssertion: aws.String("ASSERT"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}

func TestUnsignedRequest_AssumeRoleWithWebIdentity(t *testing.T) {
	req, _ := svc.AssumeRoleWithWebIdentityRequest(&sts.AssumeRoleWithWebIdentityInput{
		RoleARN:          aws.String("ARN"),
		RoleSessionName:  aws.String("SESSION"),
		WebIdentityToken: aws.String("TOKEN"),
	})

	err := req.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "", req.HTTPRequest.Header.Get("Authorization"))
}
