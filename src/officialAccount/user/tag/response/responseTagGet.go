package response

import "github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/response"

type Tag struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ResponseTagGet struct {
	response.ResponseOfficialAccount

	Tag *Tag `json:"tag"`
}

//---------------------------------------------

type ResponseTagGetList struct {
	response.ResponseOfficialAccount
	Tags []*Tag `json:"tags"`
}
