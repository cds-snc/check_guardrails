/*
Copyright © 2019 Canadian Digital Service <max.neuvians@cds-snc.ca>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	cg "github.com/cdssnc/check_guardrails/lib/aws"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AwsAudit struct {
	RootMFAEnabled     bool
	LamdbaExportExists bool
}

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Runs checks against AWS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		viper.SetDefault("output", "none")
		viper.SetDefault("aws_region", "ca-central-1")
		viper.SetDefault("breakglass_accounts", 0)
		viper.SetDefault("lambda_function", "LandingZoneLocalSNSNotificationForwarder")

		key := viper.GetString("aws_key")
		secret := viper.GetString("aws_secret")
		region := viper.GetString("aws_region")
		output := viper.GetString("output")
		lambdaFunction := viper.GetString("lambda_function")

		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(key, secret, ""),
		})

		if err != nil {
			fmt.Println("Error", err)
			return
		}

		audit := AwsAudit{}

		fmt.Println("")

		audit.RootMFAEnabled = cg.CheckRootMFA(sess)
		fmt.Println("")

		cg.CheckUserMFA(sess)
		fmt.Println("")

		cg.CheckAdminUsers(sess)
		fmt.Println("")

		audit.LamdbaExportExists = cg.CheckLambdaExport(sess, lambdaFunction)
		fmt.Println("")

		if output == "json" {
			b, _ := json.Marshal(audit)
			fmt.Println(string(b))
		}

		if !audit.RootMFAEnabled {
			os.Exit(1)
		}

		if !audit.LamdbaExportExists {
			os.Exit(1)
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)

	awsCmd.PersistentFlags().String("aws_key", "", "Your AWS key")
	awsCmd.PersistentFlags().String("aws_secret", "", "Your AWS secret")
	awsCmd.PersistentFlags().String("aws_region", "", "Your AWS region")
	awsCmd.PersistentFlags().String("lambda_function", "", "Your AWS lambda function that exports logs")
	awsCmd.PersistentFlags().String("output", "", "Output format, default: none, options: none, json")

	viper.BindPFlag("aws_key", awsCmd.PersistentFlags().Lookup("aws_key"))
	viper.BindPFlag("aws_secret", awsCmd.PersistentFlags().Lookup("aws_secret"))
	viper.BindPFlag("aws_region", awsCmd.PersistentFlags().Lookup("aws_region"))
	viper.BindPFlag("lambda_function", awsCmd.PersistentFlags().Lookup("lambda_function"))
	viper.BindPFlag("output", awsCmd.PersistentFlags().Lookup("output"))

}
