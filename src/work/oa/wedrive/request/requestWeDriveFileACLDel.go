package request

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"

type RequestWeDriveFileACLDel struct {
	UserID   string           `json:"userid"`
	FileID   string           `json:"fileid"`
	AuthInfo []*power.HashMap `json:"auth_info"`
}
