package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
)

type CompactGroupChat struct {
	ChatID string `json:"chat_id"`
	Status int    `json:"status"`
}

type ResponseGroupChatList struct {
	response.ResponseWork
	GroupChatList []*CompactGroupChat `json:"group_chat_list"`
	NextCursor    string              `json:"next_cursor"`
}

type ResponseAddJoinWay struct {
	response.ResponseWork
	ConfigId string `json:"config_id"`
}

type JoinWay struct {
	ConfigId       string   `json:"config_id"`
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark"`
	AutoCreateRoom int      `json:"auto_create_room"`
	RoomBaseName   string   `json:"room_base_name"`
	RoomBaseId     int      `json:"room_base_id"`
	ChatIdList     []string `json:"chat_id_list"`
	QrCode         string   `json:"qr_code"`
	State          string   `json:"state"`
}

type ResponseGetJoinWay struct {
	response.ResponseWork
	JoinWay JoinWay `json:"join_way"`
}

type ResponseUpdateJoinWay struct {
	response.ResponseWork
}

type ResponseDelJoinWay struct {
	response.ResponseWork
}
