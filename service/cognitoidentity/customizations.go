package cognitoidentity

import "github.com/awslabs/aws-sdk-go/aws"

func init() {
	initRequest = func(r *aws.Request) {
		switch r.Operation {
		case opGetOpenIDToken, opGetID, opGetCredentialsForIdentity:
			r.Handlers.Sign.Init() // these operations are unsigned
		}
	}
}