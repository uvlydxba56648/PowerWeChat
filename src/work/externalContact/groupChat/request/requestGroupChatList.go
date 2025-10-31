package request

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"

type RequestGroupChatList struct {
	StatusFilter int            `json:"status_filter"`
	OwnerFilter  *power.HashMap `json:"owner_filter"`
	Cursor       string         `json:"cursor"`
	Limit        int            `json:"limit"`
}

type RequestAddGroupChat struct {
	StatusFilter int            `json:"status_filter"`
	OwnerFilter  *power.HashMap `json:"owner_filter"`
	Cursor       string         `json:"cursor"`
	Limit        int            `json:"limit"`
}

type RequestGetJoinWay struct {
	ConfigId string `json:"config_id"`
}

type RequestUpdateJoinWay struct {
	ConfigId       string   `json:"config_id"`
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark"`
	AutoCreateRoom int      `json:"auto_create_room"`
	RoomBaseName   string   `json:"room_base_name"`
	RoomBaseId     int      `json:"room_base_id"`
	ChatIdList     []string `json:"chat_id_list"`
	State          string   `json:"state"`
}

type RequestDelJoinWay struct {
	ConfigId string `json:"config_id"`
}
