package request

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"

type RequestWeDriveSpaceACLDel struct {
	UserID   string           `json:"userid"`
	SpaceID  string           `json:"spaceid"`
	AuthInfo []*power.HashMap `json:"auth_info"`
}
