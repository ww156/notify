/*
# @Author ww
# @Time 2019/10/28 16:14
# @File template_test.go
*/
package template

import (
	"bytes"
	"testing"
	"text/template"
)

func TestTemplate_ExecuteHTMLString(t *testing.T) {
	var buf bytes.Buffer
	err := template.Must(template.ParseFiles("./default.tmpl")).Execute(&buf, map[string]string{"From": "!234"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(buf.String())
}
