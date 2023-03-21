package meshcmd

import (
	"log"

	"github.com/Appkube-awsx/awsx-appmesh/authenticater"
	"github.com/Appkube-awsx/awsx-appmesh/client"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/spf13/cobra"
)

var GetArnDataCmd = &cobra.Command{
	Use:   "getArnData",
	Short: "meshArnData",
	Long:  `getting Appmesh Arn Data`,
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
			getAppmeshResources(region, acKey, secKey, env)
		}
	},
}

func getAppmeshResources(region string, accessKey string, secretKey string, env string) *appmesh.ListMeshesOutput {
	log.Println("List of AWS Mesh Arn")
	appmeshClient := client.GetClient(region, accessKey, secretKey)
	appmeshResourceRequest := &appmesh.ListMeshesInput{}
	AppMeshResponse, err := appmeshClient.ListMeshes(appmeshResourceRequest)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	for _, ARN := range AppMeshResponse.Meshes {

		if env == "dev" {

			log.Println((string(*ARN.Arn)))
		}
	}

	return AppMeshResponse
}

func Execute() {
	err := GetArnDataCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}
