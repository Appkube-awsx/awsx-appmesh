package client

import (
	"github.com/Appkube-awsx/awsx-appmesh/awssession"
	//"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/service/appmesh"
	//"time"
)

func GetClient(region string, accessKey string, secretKey string) *appmesh.AppMesh {
	awsSession := awssession.GetSessionByCreds(region, accessKey, secretKey)
	svc := appmesh.New(awsSession)
	return svc
}

// func GetClient(region string, accessKey string, secretKey string) *eks.EKS {
// 	awsSession := awssession.GetSessionByCreds(region, accessKey, secretKey)
// 	svc := eks.New(awsSession)
// 	return svc}
