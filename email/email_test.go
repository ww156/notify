/*
# @Author ww
# @Time 2019/10/11 16:00
# @File email_test.go.go
*/
package email

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/ww156/notify/config"
	"os"
	"testing"
)

func TestEmail_Notify(t *testing.T) {
	e_config := config.EmailConfig{
		From:    "",
		To:      "",
		Cc:      "",
		Bcc:     "",
		Subject: "",
		Smarthost: config.HostPort{
			Host: "smtp.exmail.qq.com",
			Port: 465,
		},
		AuthUsername: "",
		AuthPassword: "",
		AuthSecret:   "",
		HTML:         "",
		Text:         "",
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	e := New(&e_config, logger)
	_, err := e.Notify(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}
