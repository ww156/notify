/*
# @Author ww
# @Time 2019/10/11 9:20
# @File notify.go
*/
package notify

import "context"

type Notifier interface {
	Notify(ctx context.Context) (bool, error)
}

// 重试通知
func Retry(noticce Notifier, count int) error {
	return nil
}
