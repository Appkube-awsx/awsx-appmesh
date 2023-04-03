package meshcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-appmesh/authenticater"
	"github.com/Appkube-awsx/awsx-appmesh/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/spf13/cobra"
)

var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "meshMetaData",
	Long:  `getting Appmesh MetaData`,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		env := cmd.Parent().PersistentFlags().Lookup("env").Value.String()

		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			meshName, _ := cmd.Flags().GetString("mesh")
			log.Println(meshName)
			getAppmesh(region, acKey, secKey, env, meshName, crossAccountRoleArn, externalId)
		}
	},
}

func getAppmesh(region string, accessKey string, secretKey string, env string, meshName string, crossAccountRoleArn string, externalId string) *appmesh.DescribeMeshOutput {
	log.Println("AWS AppMesh Metadata by MeshName")
	appmeshClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	appmeshResourceRequest := &appmesh.DescribeMeshInput{
		MeshName: aws.String(meshName),
	}
	AppMeshResponse, err := appmeshClient.DescribeMesh(appmeshResourceRequest)
	log.Println(AppMeshResponse)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return AppMeshResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("mesh", "m", "", "mesh name")

	if err := GetConfigDataCmd.MarkFlagRequired("mesh"); err != nil {
		fmt.Println("--mesh is required", err)
	}
}
