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
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kyokomi/emoji"
	"github.com/spf13/viper"

	. "github.com/logrusorgru/aurora"
)

func CheckS3Encryption(sess *session.Session, output string) bool {

	if output == "debug" {
		fmt.Println(Green("Checking AWS S3 bucket encryption settings ..."))
	}

	svc := s3.New(sess)

	result, err := svc.ListBuckets(&s3.ListBucketsInput{})

	if err != nil {
		fmt.Println("Error", err)
		return false
	}

	safelist := viper.GetStringSlice("bucket_safelist")

	set := make(map[string]bool)
	for _, v := range safelist {
		set[v] = true
	}

	for _, bucket := range result.Buckets {

		_, err := svc.GetBucketEncryption(&s3.GetBucketEncryptionInput{
			Bucket: aws.String(*bucket.Name),
		})

		if err != nil && !set[*bucket.Name] {
			if output == "debug" {
				emoji.Println(" :skull: ", BrightRed("S3 bucket found without encryption"))
				fmt.Println("")
			}
			return false
		}
	}

	if output == "debug" {
		emoji.Println(" :white_check_mark: ", BrightGreen("No unexpected S3 bucket found without encryption"))
		fmt.Println("")
	}
	return true
}
