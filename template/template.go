/*
# @Author ww
# @Time 2019/10/28 11:08
# @File template.go
*/
package template

import (
	"bytes"
	"io/ioutil"
	"os"
	tmpltext "text/template"
)

func ExecuteTextString(text string, data interface{}) (string, error) {
	if text == "" {
		return "", nil
	}
	tmpl, err := tmpltext.ParseFiles(text)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	return buf.String(), err
}

func ExecuteHTMLString(html string, data interface{}) (string, error) {
	if html == "" {
		return "", nil
	}
	tmpl, err := tmpltext.ParseFiles(html)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	return buf.String(), err
}

func ParseDingtalkMsg(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
