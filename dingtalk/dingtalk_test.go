/*
# @Author ww
# @Time 2019/10/28 17:45
# @File dingtalk_test.go
*/
package dingtalk

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/ww156/notify/config"
	"os"
	"testing"
)

func TestNotify(t *testing.T) {
	config := config.DingtalkConfig{
		AppKey:      "",
		AppSecret:   "",
		Agent_id:    0,
		Userid_list: "",
		Msg:         "",
	}
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	dd := New(&config, &logger, nil)
	dd.getToken()
	dd.Notify(context.Background())
}
