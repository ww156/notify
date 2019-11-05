/*
# @Author ww
# @Time 2019/10/28 17:33
# @File dingtalk_config.go
*/
package config

type DingtalkConfig struct {
	AppKey       string      `json:"app_key"`
	AppSecret    string      `json:"app_secret"`
	From         string      `json:"from"`
	To           string      `json:"to"`
	Agent_id     uint        `json:"agent_id"`
	Userid_list  string      `json:"userid_list"`
	Dept_id_list string      `json:"dept_id_list"`
	To_all_user  bool        `json:"to_all_user"`
	Msg          string      `json:"msg"`  // 消息模板文件
	Data         interface{} `json:"data"` // 消息数据
}
