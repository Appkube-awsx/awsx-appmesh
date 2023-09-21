package appmeshcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
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

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		
		if err != nil {
			cmd.Help()
			return
		}
		
		if authFlag {
			meshName, _ := cmd.Flags().GetString("meshName")

			if meshName != "" {
				GetMeshDetails(meshName, *clientAuth)
			} else {
				log.Fatalln("meshName not provided.Program exit")
			}

		}
	},
}


func GetMeshDetails(meshName string,  auth client.Auth) (*appmesh.DescribeMeshOutput, error) {
	log.Println("Getting aws mesh data")

	listMeshClient := client.GetClient(auth, client.APPMESH_CLIENT).(*appmesh.AppMesh)

	input := &appmesh.DescribeMeshInput{
		MeshName: aws.String(meshName),
	}
	
	meshesDetailsResponse, err := listMeshClient.DescribeMesh(input)

	
	if err != nil { 
		log.Fatalln("Error:", err)
	}

	log.Println(meshesDetailsResponse)
	return meshesDetailsResponse,err

}
	
func init() {
	GetConfigDataCmd.Flags().StringP("meshName", "t", "", "mesh Name ")

	if err := GetConfigDataCmd.MarkFlagRequired("meshName"); err != nil {
		fmt.Println(err)
	}
	
}
