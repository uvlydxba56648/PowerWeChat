package request

// JumpWxa 跳转小程序配置
type JumpWxa struct {
	Path       string `json:"path,omitempty"`        // 已发布的小程序页面路径
	Query      string `json:"query,omitempty"`       // 页面参数，最大1024字符
	EnvVersion string `json:"env_version,omitempty"` // release/trial/develop
}

// NFCUrlLinkGenerate NFC URL Link 生成请求
type NFCUrlLinkGenerate struct {
	ModelID string   `json:"model_id"`           // 必填：设备型号ID
	SN      string   `json:"sn,omitempty"`       // 选填：设备序列号
	JumpWxa *JumpWxa `json:"jump_wxa,omitempty"` // 选填：跳转配置
}
