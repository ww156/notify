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

// 重试
func Retry(notice Notifier, count int) error {
	ok, err := notice.Notify(context.Background())
	if ok {
		return nil
	}
	for i := 0; i < count; i++ {
		ok, err = notice.Notify(context.Background())
		if ok {
			return nil
		}
	}
	if err != nil {
		return err
	}
	return nil
}
