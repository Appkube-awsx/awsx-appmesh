package cmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-appmesh/client"
	"github.com/Appkube-awsx/awsx-appmesh/vault"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var awsxServiceMeshCmd = &cobra.Command{
	Use:   "getElementDetails",
	Short: "getElementDetails command gets resource counts",
	Long:  `getElementDetails command gets resource counts details of an AWS account`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getAppMesh.metadata started")
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

func getAppmeshResources(region string, accessKey string, secretKey string, env string) *appmesh.DescribeMeshOutput {
	log.Println("AWS AppMesh metadata by Mesh")
	appmeshClient := client.GetClient(region, accessKey, secretKey)
	appmeshResourceRequest := &appmesh.DescribeMeshInput{
		// mesh name - (abdul-test-1)
		MeshName: aws.String("abdul-test-1"),
	}
	AppMeshResponse, err := appmeshClient.DescribeMesh(appmeshResourceRequest)

	log.Println(AppMeshResponse)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	return AppMeshResponse
}

func Execute() {
	err := awsxServiceMeshCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	awsxServiceMeshCmd.Flags().String("vaultUrl", "", "vault end point")
	awsxServiceMeshCmd.Flags().String("accountId", "", "aws account number")
	awsxServiceMeshCmd.Flags().String("zone", "", "aws region")
	awsxServiceMeshCmd.Flags().String("accessKey", "", "aws access key")
	awsxServiceMeshCmd.Flags().String("secretKey", "", "aws secret key")
	awsxServiceMeshCmd.Flags().String("env", "", "aws env Resquired")
	//awsxServiceMeshCmd.Flags().String("getMetaData", "", "aws env Resquired")
	//AwsxCloudElementsCmd.Flags().String("crossAccountRoleArn", "", "aws cross account role arn")
}
