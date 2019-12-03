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
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/cdssnc/check_guardrails/lib/i18n"

	"github.com/kyokomi/emoji"
	. "github.com/logrusorgru/aurora"
)

func CheckLambdaExport(sess *session.Session, needle, output string) bool {

	if output == "debug" {
		fmt.Println(Green(i18n.T("check_lambda")))
	}

	svc := lambda.New(sess)

	result, err := svc.ListFunctions(&lambda.ListFunctionsInput{})

	if err != nil {
		fmt.Println("Error", err)
		return false
	}

	found := false

	for _, lf := range result.Functions {
		if lf == nil {
			continue
		}
		if *lf.FunctionName == needle {
			found = true
		}
	}

	if found {
		if output == "debug" {
			emoji.Println(" :white_check_mark: ", BrightGreen(i18n.T("check_lambda_pass")))
			fmt.Println("")
		}
		return true
	} else {
		if output == "debug" {
			emoji.Println(" :x: ", Bold(BrightRed(i18n.T("check_lambda_fail"))))
			fmt.Println("")
		}
		return false
	}

}
