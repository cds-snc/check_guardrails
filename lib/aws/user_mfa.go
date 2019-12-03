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
package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/cdssnc/check_guardrails/lib/i18n"

	"github.com/kyokomi/emoji"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
)

func CheckUserMFA(sess *session.Session, output string) bool {

	if output == "debug" {
		fmt.Println(Green(i18n.T("check_user_mfa")))
	}

	svc := iam.New(sess)

	result, err := svc.ListUsers(&iam.ListUsersInput{})

	if err != nil {
		fmt.Println("Error", err)
		return false
	}

	mfaAccounts := 0
	consoleUsers := 0
	breakglassAccounts := viper.GetInt("breakglass_accounts")

	for _, user := range result.Users {
		if user == nil {
			continue
		}

		password, _ := svc.GetLoginProfile(&iam.GetLoginProfileInput{
			UserName: user.UserName,
		})

		if password.LoginProfile != nil {
			consoleUsers++
			devices, err := svc.ListMFADevices(&iam.ListMFADevicesInput{
				UserName: user.UserName,
			})
			if err != nil {
				fmt.Println("Error", err)
				return false
			}
			if len(devices.MFADevices) != 0 {
				mfaAccounts++
			}
		}
	}

	if (consoleUsers - breakglassAccounts) != mfaAccounts {
		if output == "debug" {
			emoji.Println(" :x: ", Sprintf(BrightRed(i18n.T("check_user_mfa_fail")), mfaAccounts, consoleUsers))
			fmt.Println("")
		}
		return false
	} else {
		if output == "debug" {
			emoji.Println(" :white_check_mark: ", Sprintf(BrightGreen(i18n.T("check_user_mfa_pass")), breakglassAccounts))
			fmt.Println("")
		}
		return true
	}
}
