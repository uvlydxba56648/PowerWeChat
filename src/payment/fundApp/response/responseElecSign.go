package response

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"
	"time"
)

type ResponseApplyForElecSign struct {
	response.ResponsePayment

	State      string    `json:"state"`
	CreateTime time.Time `json:"create_time"`
}

type ResponseQueryElecSign struct {
	State       string    `json:"state"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	HashType    string    `json:"hash_type"`
	HashValue   string    `json:"hash_value"`
	DownloadUrl string    `json:"download_url"`
}

type RequestDownloadUrl struct {
	HashType    string `json:"hash_type"`
	HashValue   string `json:"hash_value"`
	DownloadUrl string `json:"download_url"`
}
