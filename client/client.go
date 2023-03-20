package client

import (
	"github.com/Appkube-awsx/awsx-appmesh/awssession"
	//"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/service/appmesh"
)

func GetClient(region string, accessKey string, secretKey string) *appmesh.AppMesh {
	awsSession := awssession.GetSessionByCreds(region, accessKey, secretKey)
	svc := appmesh.New(awsSession)
	return svc
}
