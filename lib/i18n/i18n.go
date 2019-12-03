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
package i18n

import "github.com/spf13/viper"

type Dict struct {
	en, fr string
}

func Dicts() map[string]Dict {
	m := make(map[string]Dict)

	m["check_root_mfa"] = Dict{
		"Checking AWS root account for MFA ...",
		"Vérification du compte racine AWS pour l'AMF ....",
	}

	return m
}

func T(key string) string {
	dict := Dicts()

	if _, ok := dict[key]; !ok {
		return key
	}

	if viper.GetString("loc") == "fr" {
		return dict[key].fr
	} else {
		return dict[key].en
	}
}
