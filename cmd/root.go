package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"

	"github.com/Appkube-awsx/awsx-appmesh/cmd/appmeshcmd"

	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/spf13/cobra"
)

// AwsxMeshesMetadataCmd represents the base command when called without any subcommands

var AwsxMeshesMetadataCmd = &cobra.Command{
	Use:   "getListMeshesMetaDataDetails",
	Short: "getListMeshesMetaDataDetails command gets resource counts",
	Long:  `getListMeshesMetaDataDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command meshes started")

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			GetListCluster(*clientAuth)
		} else {
			cmd.Help()
			return
		}
	},
}


// json.Unmarshal
func GetListCluster(auth client.Auth) (*appmesh.ListMeshesOutput, error) {

	log.Println("getting appmeshes list summary")

	listMeshClient := client.GetClient(auth, client.APPMESH_CLIENT).(*appmesh.AppMesh)

	listMeshRequest := &appmesh.ListMeshesInput{}
	
	listMeshResponse, err := listMeshClient.ListMeshes(listMeshRequest)

	if err != nil {
		log.Fatalln("Error:in getting  meshes list", err)
	}
 
	log.Println(listMeshResponse)

	return listMeshResponse, err
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
	AwsxMeshesMetadataCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxMeshesMetadataCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxMeshesMetadataCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxMeshesMetadataCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxMeshesMetadataCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxMeshesMetadataCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxMeshesMetadataCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
