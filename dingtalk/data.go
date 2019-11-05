/*
# @Author ww
# @Time 2019/10/28 17:55
# @File data.go
*/
package dingtalk

type Res_token struct {
	Errcode      int
	Access_token string
	Errmsg       string
	Expires_in   int
}

type Msg struct {
	Msgtype string `json:"msgtype"`
	// 文本消息
	Text *struct {
		Content string `json:"content"`
	} `json:"text,omitempty"`
	// 图片消息
	Image *struct {
		Media_id string `json:"media_id"`
	} `json:"image,omitempty"`
	// 语音消息
	Voice *struct {
		Medis_id string `json:"medis_id"`
		Duration string `json:"duration"`
	} `json:"voice,,omitempty"`
	// 文件消息
	File *struct {
		//媒体文件id。引用的媒体文件最大10MB
		Media_id string `json:"media_id"`
	} `json:"file,omitempty"`
	// 链接消息
	Link *struct {
		Title      string `json:"title"`
		Text       string `json:"text"`
		PicUrl     string `json:"pic_url"`
		MessageUrl string `json:"message_url"`
	} `json:"link,omitempty"`
	// OA消息
	Oa *struct {
		Message_url string
		Head        struct {
			Bgcolor string
			Text    string
		}
		Body struct {
			Title string
			Form  []struct {
				Key   string
				Value string
			}
			Rich struct {
				Num  string
				Uint string
			}
			Content    string
			Image      string
			File_count string
			Author     string
		}
	} `json:"oa,omitempty"`
	// markdown消息
	Markdown *struct {
		Title string
		Text  string
	} `json:"markdown,omitempty"`
	// 卡片消息
	Action_card *struct {
		Title           string
		Markdown        string
		Single_title    string
		Single_url      string
		Btn_orientation string
		Btn_json_list   []struct {
			Title      string
			Action_url string
		}
	} `json:"action_card,omitempty"`
}
