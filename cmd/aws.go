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
	"fmt"

	"github.com/cdssnc/check_guardrails/lib/aws"
	"github.com/spf13/cobra"
)

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Runs checks against AWS",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString("aws_key")

		if err != nil {
			fmt.Println("Error", err)
			return
		}

		secret, err := cmd.Flags().GetString("aws_secret")

		if err != nil {
			fmt.Println("Error", err)
			return
		}
		fmt.Println("")
		aws.CheckRootMFA(key, secret)
		fmt.Println("")
		aws.CheckUserMFA(key, secret)
		fmt.Println("")
		aws.CheckAdminUsers(key, secret)
		fmt.Println("")
		aws.CheckLambdaExport(key, secret, "LandingZoneLocalSNSNotificationForwarder")
		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(awsCmd)

	awsCmd.PersistentFlags().String("aws_key", "", "Your AWS key")
	awsCmd.PersistentFlags().String("aws_secret", "", "Your AWS secret")
	awsCmd.PersistentFlags().String("lambda_function", "", "Your AWS lambda function that exports logs")

}
