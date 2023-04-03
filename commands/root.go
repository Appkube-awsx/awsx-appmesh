package commands

import (
	"log"

	"github.com/Appkube-awsx/awsx-appmesh/authenticater"
	"github.com/Appkube-awsx/awsx-appmesh/client"
	"github.com/Appkube-awsx/awsx-appmesh/commands/meshcmd"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxServiceMeshCmd = &cobra.Command{
	Use:   "GetAppMeshList",
	Short: "GetAppMeshList command gets resource Arn",
	Long:  `GetAppMeshList command gets resource Arn details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getAppmeshList started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		env := cmd.PersistentFlags().Lookup("env").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			getAppmeshResource(region, acKey, secKey, env, crossAccountRoleArn, externalId)
		}
	},
}

type Tags struct {
	Environment string `json:"Environment"`
}

func getAppmeshResource(region string, accessKey string, secretKey string, env string, crossAccountRoleArn string, externalId string) *appmesh.ListMeshesOutput {
	log.Println("List of AWS Mesh")
	appmeshClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	appmeshResourceRequest := &appmesh.ListMeshesInput{}
	AppMeshResponse, err := appmeshClient.ListMeshes(appmeshResourceRequest)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	for _, List := range AppMeshResponse.Meshes {

		if env == "dev" {
			log.Println(List)

		}
	}

	return AppMeshResponse
}

func Execute() {
	err := AwsxServiceMeshCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxServiceMeshCmd.AddCommand(meshcmd.GetConfigDataCmd)
	AwsxServiceMeshCmd.AddCommand(meshcmd.GetArnDataCmd)
	AwsxServiceMeshCmd.AddCommand(meshcmd.GetCostDataCmd)
	AwsxServiceMeshCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxServiceMeshCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxServiceMeshCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxServiceMeshCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxServiceMeshCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxServiceMeshCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxServiceMeshCmd.PersistentFlags().String("externalId", "", "aws external id auth")
	AwsxServiceMeshCmd.PersistentFlags().String("env", "", "env")

}
