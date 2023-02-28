package cmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-appmesh/client"
	"github.com/Appkube-awsx/awsx-appmesh/vault"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxServiceMeshCmd = &cobra.Command{
	Use:   "getElementDetails",
	Short: "getElementDetails command gets resource counts",
	Long:  `getElementDetails command gets resource counts details of an AWS account`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getAppmeshArn started")
		vaultUrl, _ := cmd.Flags().GetString("vaultUrl")
		accountNo, _ := cmd.Flags().GetString("accountId")
		region, _ := cmd.Flags().GetString("zone")
		acKey, _ := cmd.Flags().GetString("accessKey")
		secKey, _ := cmd.Flags().GetString("secretKey")
		env, _ := cmd.Flags().GetString("env")
		//crossAccountRoleArn, _ := cmd.Flags().GetString("crossAccountRoleArn")

		if vaultUrl != "" && accountNo != "" && env != "" {
			fmt.Println("in vault")
			if region == "" {
				log.Fatalln("Zone not provided. Program exit")
				return
			}
			log.Println("Getting account details")
			data, err := vault.GetAccountDetails(vaultUrl, accountNo)
			if err != nil {
				log.Println("Error in calling the account details api. \n", err)
				return
			}
			if data.AccessKey == "" || data.SecretKey == "" {
				log.Println("Account details not found.")
				return
			}
			getAppmeshResources(region, data.AccessKey, data.SecretKey, env)
		} else if region != "" && acKey != "" && secKey != "" && env != "" {
			getAppmeshResources(region, acKey, secKey, env)
		} else {
			log.Fatal("region", secKey)
			log.Fatal("AWS credentials like accesskey/secretkey/region/crossAccountRoleArn not provided. Program exit")
			return
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
			//log.Println(xyz)
			log.Println((string(*ARN.Arn)))
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
	AwsxServiceMeshCmd.Flags().String("vaultUrl", "", "vault end point")
	AwsxServiceMeshCmd.Flags().String("accountId", "", "aws account number")
	AwsxServiceMeshCmd.Flags().String("zone", "", "aws region")
	AwsxServiceMeshCmd.Flags().String("accessKey", "", "aws access key")
	AwsxServiceMeshCmd.Flags().String("secretKey", "", "aws secret key")
	AwsxServiceMeshCmd.Flags().String("env", "", "aws env Resquired")
	//AwsxCloudElementsCmd.Flags().String("crossAccountRoleArn", "", "aws cross account role arn")
}
