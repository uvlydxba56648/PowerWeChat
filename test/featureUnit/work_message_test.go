package featureUnit

import (
	"context"
	"testing"

	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
)

func Test_Message_Send_Text(t *testing.T) {
	//para := &request.RequestMessageSendText{
	//	&request.RequestMessageSend{
	//		ToUser:                 "michael",
	//		MsgType:                "text",
	//		AgentID:                1000005,
	//		Safe:                   0,
	//		EnableIDTrans:          0,
	//		EnableDuplicateCheck:   0,
	//		DuplicateCheckInterval: 1800,
	//	},
	//	&request.RequestText{
	//		Content: "您已经成功测试收到系统推送消息",
	//	},
	//}

	powerPara := &power.HashMap{
		"touser":  "michaelhu",
		"msgtype": "text",
		"agentid": 1000004,
		"text": power.HashMap{
			"content": "您已经成功测试收到系统推送消息",
		},
		"safe":                     0,
		"enable_id_trans":          0,
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	}
	response, err := Work.Message.Send(context.Background(), powerPara)
	if err != nil {
		t.Error("uniformMessage send err is ", err)
	}

	if response == nil && response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMsg)
	}
}
