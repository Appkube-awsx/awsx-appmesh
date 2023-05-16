package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-appmesh/authenticator"
	"github.com/Appkube-awsx/awsx-appmesh/client"

	"github.com/Appkube-awsx/awsx-appmesh/cmd/appmeshcmd"

	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/spf13/cobra"
)

var AwsxMeshesMetadataCmd = &cobra.Command{
	Use:   "getListMeshesMetaDataDetails",
	Short: "getListMeshesMetaDataDetails command gets resource counts",
	Long:  `getListMeshesMetaDataDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command meshes started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			getListCluster(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}


// json.Unmarshal
func getListCluster(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*appmesh.ListMeshesOutput) {
	log.Println("getting meshes metadata list summary")

	listClusterClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	listClusterRequest := &appmesh.ListMeshesInput{}
	
	listClusterResponse, err := listClusterClient.ListMeshes(listClusterRequest)
	if err != nil {
		log.Fatalln("Error:in getting  meshes list", err)
	}
 
	log.Println(listClusterResponse)
	return listClusterResponse
}

func Execute() {
	err := AwsxMeshesMetadataCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxMeshesMetadataCmd.AddCommand(appmeshcmd.GetConfigDataCmd)

	AwsxMeshesMetadataCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxMeshesMetadataCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxMeshesMetadataCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxMeshesMetadataCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxMeshesMetadataCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxMeshesMetadataCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxMeshesMetadataCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
