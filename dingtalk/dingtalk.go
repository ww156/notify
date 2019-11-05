/*
# @Author ww
# @Time 2019/10/28 17:45
# @File dingtalk.go
*/
package dingtalk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"notify/config"
	"notify/template"
	"time"
)

const (
	// access_token
	OAPI = "https://oapi.dingtalk.com/gettoken"
	// 工作通知消息
	MSG_NOTIFY_URL = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2"
	// 群消息
	MSG_CHAT_URL = "https://oapi.dingtalk.com/chat/send"
	// 普通消息
	MSG_URL = "https://oapi.dingtalk.com/message/send_to_conversation"
)

type Dingtalk struct {
	ctx    context.Context
	conf   *config.DingtalkConfig
	l      *log.Logger
	token  string
	client *http.Client
}

func New(conf *config.DingtalkConfig, l *log.Logger, client *http.Client) *Dingtalk {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*7200)
	c := client
	if c == nil {
		c = http.DefaultClient
	}
	token, err := get_token(c, conf.AppKey, conf.AppSecret)
	if err != nil {
		for i := 0; i < 3 && err != nil; i++ {
			token, err = get_token(client, conf.AppKey, conf.AppSecret)
		}
		if err != nil {
			panic(err)
		}
	}

	return &Dingtalk{
		ctx:    ctx,
		conf:   conf,
		l:      l,
		token:  token,
		client: c,
	}
}

// manager3847
func (d *Dingtalk) Notify(ctx context.Context) (bool, error) {
	_, err := d.work_notice()
	if err != nil {
		fmt.Println(err)
	}
	return true, nil
}

func (d *Dingtalk) GetToken() error {
	go func() {
		conf := d.conf
		select {
		case <-d.ctx.Done():
			token, err := get_token(d.client, conf.AppKey, conf.AppSecret)
			if err != nil {
				fmt.Println(err)
			}
			d.token = token
		}
	}()
	return nil
}

// 根据手机号获取userid
func (d *Dingtalk) get_by_mobile(mobile string) (string, error) {
	return "", nil
}

// 工作通知
func (d *Dingtalk) work_notice() (string, error) {
	conf := d.conf
	v := url.Values{}
	v.Set("access_token", d.token)
	b, err := template.ParseDingtalkMsg(conf.Msg)
	if err != nil {
		return "", err
	}
	var msg Msg
	err = json.Unmarshal(b, &msg)
	if err != nil {
		return "", err
	}
	data := req_work_notice{
		Agent_id:    conf.Agent_id,
		Userid_list: conf.Userid_list,
		Msg:         msg,
	}
	m, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", MSG_NOTIFY_URL+"?"+v.Encode(), bytes.NewReader(m))
	if err != nil {
		return "", err
	}
	res, err := d.client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	res.Body.Close()
	r := Res_work_notice{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", nil
	}
	fmt.Println(r, d.token, string(m))
	if r.Errcode != 0 {
		return "", fmt.Errorf("errcode %d", r.Errcode)
	}
	return "", nil
}

// 获取token
func get_token(client *http.Client, appkey, appsecret string) (string, error) {
	v := url.Values{}
	v.Set("appkey", appkey)
	v.Set("appsecret", appsecret)
	req, err := http.NewRequest("GET", OAPI+"?"+v.Encode(), nil)
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	res.Body.Close()
	token := Res_token{}
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", nil
	}
	if token.Errcode != 0 {
		return "", fmt.Errorf("errcode %d", token.Errcode)
	}
	return token.Access_token, nil
}

type req_work_notice struct {
	Agent_id    uint   `json:"agent_id"`
	Userid_list string `json:"userid_list"`
	Msg         Msg    `json:"msg"`
}

type Res_work_notice struct {
	Errcode int
	Task_id int
	Request string
}
