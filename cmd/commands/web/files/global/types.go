package global

type M map[string]interface{}

type MStr map[string]string

type JsonResp struct {
	ErrCode int64       `json:"errCode" default:"0"` // 错误码
	Data    interface{} `json:"data,omitempty"`      // 返回数据
	Msg     interface{} `json:"msg,omitempty"`       // 提示信息/错误信息
	Debug   interface{} `json:"debug,omitempty"`     // 调试数据,仅在非生产环境出现
}
