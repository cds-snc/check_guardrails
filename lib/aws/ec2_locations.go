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
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/kyokomi/emoji"

	. "github.com/logrusorgru/aurora"
)

func CheckEC2Residency(sess *session.Session, regions *ec2.DescribeRegionsOutput, output string) bool {

	if output == "debug" {
		fmt.Println(Green("Checking AWS EC2 data residency ..."))
	}

	for _, region := range regions.Regions {

		newSess := sess.Copy(&aws.Config{Region: aws.String(*region.RegionName)})
		newSvc := ec2.New(newSess)

		instances, err := newSvc.DescribeInstances(&ec2.DescribeInstancesInput{})

		if err != nil {
			fmt.Println("Error", err)
			return false
		}

		if len(instances.Reservations) > 0 && *region.RegionName != "ca-central-1" {
			if output == "debug" {
				emoji.Println(" :x: ", BrightRed("EC2 instances found outside ca-central-1"))
				fmt.Println("")
			}
			return false
		}
	}

	if output == "debug" {
		emoji.Println(" :white_check_mark: ", BrightGreen("No EC2 instances found outside ca-central-1"))
		fmt.Println("")
	}
	return true
}
