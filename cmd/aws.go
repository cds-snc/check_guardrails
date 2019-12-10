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
	"github.com/aws/aws-sdk-go/service/ec2"

	cg "github.com/cds-snc/check_guardrails/lib/aws"
	"github.com/cds-snc/check_guardrails/lib/html"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AwsAudit struct {
	RootMFAEnabled     bool
	NoRootKeys         bool
	UsersHaveMFA       bool
	NoAdminUsers       bool
	LamdbaExportExists bool
	StrongPasswords    bool
	GuardDutyActive    bool
	EC2Residency       bool
	S3Encrypted        bool
	RDSEncrypted       bool
	NoPort80           bool
}

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Runs checks against AWS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

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

		regionSvc := ec2.New(sess)
		regions, err := regionSvc.DescribeRegions(&ec2.DescribeRegionsInput{})

		if err != nil {
			fmt.Println("Error", err)
			return
		}

		audit := AwsAudit{}

		if output == "debug" {
			fmt.Println("")
		}

		audit.RootMFAEnabled = cg.CheckRootMFA(sess, output)
		audit.NoRootKeys = cg.CheckRootKeys(sess, output)
		audit.UsersHaveMFA = cg.CheckUserMFA(sess, output)
		audit.NoAdminUsers = cg.CheckAdminUsers(sess, output)
		audit.LamdbaExportExists = cg.CheckLambdaExport(sess, lambdaFunction, output)
		audit.StrongPasswords = cg.CheckPasswordPolicy(sess, output)
		// cg.CheckSSO(sess, output)
		audit.GuardDutyActive = cg.CheckGuardDuty(sess, output)
		audit.EC2Residency = cg.CheckEC2Residency(sess, regions, output)
		audit.S3Encrypted = cg.CheckS3Encryption(sess, output)
		audit.RDSEncrypted = cg.CheckRDSEncryption(sess, regions, output)
		audit.NoPort80 = cg.CheckSecurityGroupsPort80(sess, regions, output)

		if output == "json" {
			b, _ := json.Marshal(audit)
			fmt.Println(string(b))
		}

		if output == "html" {
			html.Render(audit)
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
	awsCmd.PersistentFlags().String("breakglass_accounts", "", "Number of break glass accounts to expect")
	awsCmd.PersistentFlags().String("output", "", "Output format, default: none, options: none, json")

	viper.SetDefault("output", "debug")
	viper.SetDefault("aws_region", "ca-central-1")
	viper.SetDefault("breakglass_accounts", 0)
	viper.SetDefault("lambda_function", "LandingZoneLocalSNSNotificationForwarder")
	viper.SetDefault("bucket_safelist", []string{})

	viper.BindPFlag("aws_key", awsCmd.PersistentFlags().Lookup("aws_key"))
	viper.BindPFlag("aws_secret", awsCmd.PersistentFlags().Lookup("aws_secret"))
	viper.BindPFlag("aws_region", awsCmd.PersistentFlags().Lookup("aws_region"))
	viper.BindPFlag("lambda_function", awsCmd.PersistentFlags().Lookup("lambda_function"))
	viper.BindPFlag("breakglass_accounts", awsCmd.PersistentFlags().Lookup("breakglass_accounts"))
	viper.BindPFlag("output", awsCmd.PersistentFlags().Lookup("output"))

}
