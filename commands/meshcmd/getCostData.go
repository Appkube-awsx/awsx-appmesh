package meshcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-appmesh/authenticater"
	"github.com/Appkube-awsx/awsx-appmesh/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var GetCostDataCmd = &cobra.Command{
	Use:   "getCostData",
	Short: "meshCostData",
	Long:  `getting Appmesh Cost and Usage Data`,
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
			getAppmeshCost(region, crossAccountRoleArn, acKey, secKey, env, externalId)
		}
	},
}

type Tags struct {
	Environment string `json:"Environment"`
}

func getAppmeshCost(region string, crossAccountRoleArn string, accessKey string, secretKey string, env string, externalId string) *costexplorer.GetCostAndUsageOutput {
	Client := client.GetCostClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	// costExplorer := costexplorer.(Client)
	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String("2023-02-01"),
			End:   aws.String("2023-02-28"),
		},
		Granularity: aws.String("MONTHLY"),
		Metrics: []*string{
			aws.String("USAGE_QUANTITY"),
			aws.String("UNBLENDED_COST"),
			aws.String("BLENDED_COST"),
			aws.String("AMORTIZED_COST"),
			aws.String("NET_AMORTIZED_COST"),
			aws.String("NET_UNBLENDED_COST"),
			aws.String("NORMALIZED_USAGE_AMOUNT"),
		},
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("RECORD_TYPE"),
			},
		},

		Filter: &costexplorer.Expression{
			And: []*costexplorer.Expression{
				{
					Dimensions: &costexplorer.DimensionValues{
						Key: aws.String("SERVICE"),
						Values: []*string{
							aws.String("Amazon App Mesh"),
						},
					},
				},
				{
					Dimensions: &costexplorer.DimensionValues{
						Key: aws.String("RECORD_TYPE"),
						Values: []*string{
							aws.String("Credit"),
						},
					},
				},
			},
		},
	}

	result, err := Client.GetCostAndUsage(input)
	log.Println(result)

	if err != nil {
		fmt.Println("Error getting cost and usage:", err)
		return result
	}

	return result
}
