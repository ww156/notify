/*
# @Author ww
# @Time 2019/11/5 14:23
# @File sms.go
*/
package sms

import (
	"context"
	"log"
	"notify/config"
)

type Sms struct {
	conf   *config.SmsConfig
	logger *log.Logger
}

func New(c *config.SmsConfig, l *log.Logger) *Sms {
	return &Sms{conf: c, logger: l}
}

func (s *Sms) Notify(ctx context.Context) (bool, error) {
	return true, nil
}
