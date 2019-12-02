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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/kyokomi/emoji"

	. "github.com/logrusorgru/aurora"
)

func CheckGuardDuty(sess *session.Session, output string) bool {

	if output == "debug" {
		fmt.Println(Green("Checking AWS GuardDuty ..."))
	}

	svc := guardduty.New(sess)

	result, err := svc.ListDetectors(&guardduty.ListDetectorsInput{})

	if err != nil {
		fmt.Println("Error", err)
		return false
	}

	if len(result.DetectorIds) == 0 {
		if output == "debug" {
			emoji.Println(" :x: ", BrightRed("No GuardDuty detectors found!"))
			fmt.Println("")
		}
		return false
	}

	detector := result.DetectorIds[0]

	accountResult, err := svc.GetMasterAccount(&guardduty.GetMasterAccountInput{
		DetectorId: aws.String(*detector),
	})

	if err != nil {
		fmt.Println("Error", err)
		return false
	}

	if accountResult.Master != nil && accountResult.Master.RelationshipStatus != nil {
		if *accountResult.Master.RelationshipStatus == "Enabled" {
			if output == "debug" {
				emoji.Println(" :white_check_mark: ", BrightGreen("GuardDuty found with master account enabled"))
				fmt.Println("")
			}
			return true
		}
	}

	if output == "debug" {
		emoji.Println(" :x: ", BrightRed("No GuardDuty master account found or not enabled"))
		fmt.Println("")
	}
	return false
}
