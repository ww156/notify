/*
# @Author ww
# @Time 2019/10/11 9:20
# @File notify.go
*/
package notify

import "context"

type Mode string

type Notifier interface {
	Notify(ctx context.Context) (bool, error)
}
