package controller

import (
	"github.com/Appkube-awsx/awsx-appmesh/cmd"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"log"
)

func GetAppmeshByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*appmesh.ListMeshesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetAppmeshByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetAppmeshByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) (*appmesh.ListMeshesOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetAppmeshByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetAppmeshByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*appmesh.ListMeshesOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := cmd.GetListCluster(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetAppmesh(clientAuth *client.Auth) (*appmesh.ListMeshesOutput, error) {
	response, err := cmd.GetListCluster(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
