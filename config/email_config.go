/*
# @Author ww
# @Time 2019/10/15 13:54
# @File email_config.go
*/
package config

import (
	"strconv"
)

const secretToken = "<secret>"

type Secret string

type HostPort struct {
	Host string
	Port int
}

func (h *HostPort) String() string {
	return h.Host + ":" + strconv.Itoa(h.Port)
}

// EmailConfig configures notifications via mail.
type EmailConfig struct {
	// Email address to notify.
	From         string      `yaml:"from,omitempty" json:"from,omitempty"`
	To           string      `yaml:"to,omitempty" json:"to,omitempty"`
	Cc           string      `yaml:"cc,omitempty" json:"cc,omitempty"`
	Bcc          string      `yaml:"bcc,omitempty" json:"bcc,omitempty"`
	Subject      string      `yaml:"Subject,omitempty" json:"subject"`
	Smarthost    HostPort    `yaml:"smarthost,omitempty" json:"smarthost,omitempty"`
	AuthUsername string      `yaml:"auth_username,omitempty" json:"auth_username,omitempty"`
	AuthPassword Secret      `yaml:"auth_password,omitempty" json:"auth_password,omitempty"`
	AuthSecret   Secret      `yaml:"auth_secret,omitempty" json:"auth_secret,omitempty"`
	HTML         string      `yaml:"html,omitempty" json:"html,omitempty"`
	Text         string      `yaml:"text,omitempty" json:"text,omitempty"`
	Data         interface{} `yaml:"data" json:"data"`
}
