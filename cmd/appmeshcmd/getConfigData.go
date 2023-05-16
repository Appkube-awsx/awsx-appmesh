package appmeshcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-appmesh/authenticator"
	"github.com/Appkube-awsx/awsx-appmesh/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		// print(authFlag)
		// authFlag := true
		
		if authFlag {
			meshName, _ := cmd.Flags().GetString("meshName")
			meshOwner, _ := cmd.Flags().GetString("meshOwner")
			
			
			if meshName != "" {
				getMeshDetails(region, crossAccountRoleArn,  acKey, secKey, meshName, meshOwner, externalId)

			} else {
				log.Fatalln("meshName not provided.Program exit")
			}
			if meshOwner != "" {
				getMeshDetails(region, crossAccountRoleArn,  acKey, secKey, meshName, meshOwner,  externalId)

			} else {
				log.Fatalln("meshOwner not provided.Program exit")
			}
			

		}
	},
}


func getMeshDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, meshName string, meshOwner string, externalId string) *appmesh.DescribeMeshOutput {
	log.Println("Getting aws mesh data")

	listClusterClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	input := &appmesh.DescribeMeshInput{
		MeshName: aws.String(meshName),
		MeshOwner: aws.String(meshOwner),
	}
	
	clusterDetailsResponse, err := listClusterClient.DescribeMesh(input)

	
	if err != nil { 
		log.Fatalln("Error:", err)
	}

	log.Println(clusterDetailsResponse)
	return clusterDetailsResponse

}
	
func init() {
	GetConfigDataCmd.Flags().StringP("meshName", "t", "", "mesh Name ")

	if err := GetConfigDataCmd.MarkFlagRequired("meshName"); err != nil {
		fmt.Println(err)
	}

	GetConfigDataCmd.Flags().StringP("meshOwner", "u", "", "mesh owner name")

	if err := GetConfigDataCmd.MarkFlagRequired("meshOwner"); err != nil {
		fmt.Println(err)
	}
}
