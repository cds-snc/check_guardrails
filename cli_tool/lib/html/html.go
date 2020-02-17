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
package html

import (
	"fmt"
	"html/template"
	"os"
)

func Render(audit interface{}) {
	t := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Check guardrails report</title>
		<meta name="author" content="name">
		<meta name="description" content="description here">
		<meta name="keywords" content="keywords,here">
		<link href="https://cdn.jsdelivr.net/npm/tailwindcss/dist/tailwind.min.css" rel="stylesheet"> 
	</head>
	<body class="font-sans leading-normal tracking-normal">
		<div class="container w-full md:max-w-xl mx-auto pt-20">
			<div class="w-full px-4 md:px-6 text-xl text-grey-darkest leading-normal bg-white rounded border-2" style="font-family:Georgia,serif;">
				<div class="font-sans pb-6">
                    <h1 class="font-sans break-normal text-black pt-6 pb-2 text-2xl md:text-3xl">Check guardrails report</h1>
                    <ul class="leading-loose text-sm">
                        <li>❌  Root MFA is not enabled</li>
                        <li>✅  Root MFA has no programmatic keys</li>
                        <li>✅  All user accounts use MFA (taking into account 2 breakglass accounts)</li>
                        <li>❌  3 user(s) have admin policies attached (2 expected)</li>
                        <li>✅  Lambda export function found</li>
                        <li>✅  Password must be 15 characters or longer</li>
                        <li>✅  GuardDuty found with master account enabled</li>
                        <li>❌  EC2 instances found outside ca-central-1</li>
                        <li>✅  No unexpected S3 bucket found without encryption</li>
                        <li>❌  RDS instance found without encryption</li>
                        <li>❌  Security group with port 80 found</li>
                    </ul>
				</div>
			</div>
		</div> 
	</body>
	</html>
	`
	tmpl, err := template.New("tmpl").Parse(t)

	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(os.Stdout, audit)
}
