package models

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/contract"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/models"
)

// https://developer.work.weixin.qq.com/document/path/90967

const (
	CALLBACK_EVENT_SUBSCRIBE                = "subscribe"
	CALLBACK_EVENT_UNSUBSCRIBE              = "unsubscribe"
	CALLBACK_EVENT_ENTER_AGENT              = "enter_agent"
	CALLBACK_EVENT_LOCATION                 = "LOCATION"
	CALLBACK_EVENT_BATCH_JOB_RESULT         = "batch_job_result"
	CALLBACK_EVENT_CLICK                    = "click"
	CALLBACK_EVENT_VIEW                     = "view"
	CALLBACK_EVENT_SCANCODE_PUSH            = "scancode_push"
	CALLBACK_EVENT_SCANCODE_WAITMSG         = "scancode_waitmsg"
	CALLBACK_EVENT_PIC_SYSPHOTO             = "pic_sysphoto"
	CALLBACK_EVENT_PIC_PHOTO_OR_ALBUM       = "pic_photo_or_album"
	CALLBACK_EVENT_PIC_WEIXIN               = "pic_weixin"
	CALLBACK_EVENT_LOCATION_SELECT          = "location_select"
	CALLBACK_EVENT_OPEN_APPROVAL_CHANGE     = "open_approval_change"
	CALLBACK_EVENT_SHARE_AGENT_CHANGE       = "share_agent_change"
	CALLBACK_EVENT_TEMPLATE_CARD_EVENT      = "template_card_event"
	CALLBACK_EVENT_TEMPLATE_CARD_MENU_EVENT = "template_card_menu_event"
	CALLBACK_EVENT_SYS_APPROVAL_CHANGE      = "sys_approval_change"
	CALLBACK_EVENT_KF_MSG_OR_EVENT          = "kf_msg_or_event"
)

type EventSubscribe struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID string `xml:"AgentID" json:"AgentID"`
}

type EventEnterAgent struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey string `xml:"EventKey" json:"EventKey"`
	AgentID  string `xml:"AgentID" json:"AgentID"`
}

type EventLocation struct {
	contract.EventInterface
	models.CallbackMessageHeader
	Latitude  string `xml:"Latitude" json:"Latitude"`
	Longitude string `xml:"Longitude" json:"Longitude"`
	Precision string `xml:"Precision" json:"Precision"`
	AgentID   string `xml:"AgentID" json:"AgentID"`
	AppType   string `xml:"AppType" json:"AppType"`
}

type BatchJob struct {
	Text    string `xml:",chardata" json:",chardata"`
	JobID   string `xml:"JobId" json:"JobId"`
	JobType string `xml:"JobType" json:"JobType"`
	ErrCode string `xml:"ErrCode" json:"ErrCode"`
	ErrMsg  string `xml:"ErrMsg" json:"ErrMsg"`
}

type EventBatchJobResult struct {
	contract.EventInterface
	models.CallbackMessageHeader
	BatchJob *BatchJob `xml:"BatchJob"`
}

type EventClick struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey string `xml:"EventKey" json:"EventKey"`
	AgentID  string `xml:"AgentID" json:"AgentID"`
}

type EventView struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey string `xml:"EventKey" json:"EventKey"`
	AgentID  string `xml:"AgentID" json:"AgentID"`
}

type EventScanCodePush struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string        `xml:"EventKey" json:"EventKey"`
	ScanCodeInfo *ScanCodeInfo `xml:"ScanCodeInfo" json:"ScanCodeInfo"`
	AgentID      string        `xml:"AgentID" json:"AgentID"`
}

type ScanCodeInfo struct {
	Text       string `xml:",chardata" json:",chardata"`
	ScanType   string `xml:"ScanType" json:"ScanType"`
	ScanResult string `xml:"ScanResult" json:"ScanResult"`
}

type EventScancodeWaitMsg struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string        `xml:"EventKey" json:"EventKey"`
	ScanCodeInfo *ScanCodeInfo `xml:"ScanCodeInfo" json:"ScanCodeInfo"`
	AgentID      string        `xml:"AgentID" json:"AgentID"`
}

type EventPicSysPhoto struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string        `xml:"EventKey" json:"EventKey"`
	SendPicsInfo *SendPicsInfo `xml:"SendPicsInfo" json:"SendPicsInfo"`
	AgentID      string        `xml:"AgentID" json:"AgentID"`
}

type EventPicPhotoOrAlbum struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string        `xml:"EventKey" json:"EventKey"`
	SendPicsInfo *SendPicsInfo `xml:"SendPicsInfo" json:"SendPicsInfo"`
	AgentID      string        `xml:"AgentID" json:"AgentID"`
}

type Item struct {
	Text      string `xml:",chardata" json:",chardata"`
	PicMd5Sum string `xml:"PicMd5Sum" json:"PicMd5Sum"`
}

type PicList struct {
	Text string `xml:",chardata" json:",chardata"`
	Item *Item  `xml:"item" json:"item"`
}

type SendPicsInfo struct {
	Text    string   `xml:",chardata" json:",chardata"`
	Count   string   `xml:"Count" json:"Count"`
	PicList *PicList `xml:"PicList" json:"PicList"`
}

type EventPicWeixin struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey     string       `xml:"EventKey" json:"EventKey"`
	SendPicsInfo SendPicsInfo `xml:"SendPicsInfo" json:"SendPicsInfo"`
	AgentID      string       `xml:"AgentID" json:"AgentID"`
}

type SendLocationInfo struct {
	Text      string `xml:",chardata" json:",chardata"`
	LocationX string `xml:"Location_X" json:"Location_X"`
	LocationY string `xml:"Location_Y" json:"Location_Y"`
	Scale     string `xml:"Scale" json:"Scale"`
	Label     string `xml:"Label" json:"Label"`
	PoiName   string `xml:"Poiname" json:"Poiname"`
}

type EventLocationSelect struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey         string            `xml:"EventKey" json:"EventKey"`
	SendLocationInfo *SendLocationInfo `xml:"SendLocationInfo" json:"SendLocationInfo"`
	AgentID          string            `xml:"AgentID" json:"AgentID"`
	AppType          string            `xml:"AppType" json:"AppType"`
}

// ----------------------------------------------------------------------------

type Applier struct {
	Text   string `xml:",chardata" json:",chardata"`
	UserID string `xml:"UserId" json:"UserId"`
	Party  string `xml:"Party" json:"Party"`
}

type Approver struct {
	Text   string `xml:",chardata" json:",chardata"`
	UserID string `xml:"UserId" json:"UserId"`
}

type Detail struct {
	Text     string    `xml:",chardata" json:",chardata"`
	Approver *Approver `xml:"Approver" json:"Approver"`
	Speech   string    `xml:"Speech" json:"Speech"`
	SpStatus string    `xml:"SpStatus" json:"SpStatus"`
	SpTime   string    `xml:"SpTime" json:"SpTime"`
}

type SPRecord struct {
	Text         string   `xml:",chardata" json:",chardata"`
	SpStatus     string   `xml:"SpStatus" json:"SpStatus"`
	ApproverAttr string   `xml:"ApproverAttr" json:"ApproverAttr"`
	Details      []Detail `xml:"Details" json:"Details"`
}

type Notifier struct {
	Text   string `xml:",chardata" json:",chardata"`
	UserID string `xml:"UserId" json:"UserId"`
}

type CommentUserInfo struct {
	Text   string `xml:",chardata" json:",chardata"`
	UserID string `xml:"UserId" json:"UserId"`
}

type Comments struct {
	Text            string           `xml:",chardata" json:",chardata"`
	CommentUserInfo *CommentUserInfo `xml:"CommentUserInfo" json:"CommentUserInfo"`
	CommentTime     string           `xml:"CommentTime" json:"CommentTime"`
	CommentContent  string           `xml:"CommentContent" json:"CommentContent"`
	CommentID       string           `xml:"CommentId" json:"CommentId"`
}

type ApprovalNode struct {
	NodeStatus    int           `xml:"NodeStatus" json:"NodeStatus"`
	NodeAttr      int           `xml:"NodeAttr" json:"NodeAttr"`
	NodeType      int           `xml:"NodeType" json:"NodeType"`
	ApprovalItems ApprovalItems `xml:"Items" json:"Items"`
}

type ApprovalItems struct {
	ApprovalItem []ApprovalItem `xml:"Item" json:"Item"`
}

type ApprovalItem struct {
	ItemName   string `xml:"ItemName" json:"ItemName"`
	ItemUserId string `xml:"ItemUserId" json:"ItemUserId"`
	ItemImage  string `xml:"ItemImage" json:"ItemImage"`
	ItemStatus int    `xml:"ItemStatus" json:"ItemStatus"`
	ItemSpeech string `xml:"ItemSpeech" json:"ItemSpeech"`
	ItemOpTime int64  `xml:"ItemOpTime" json:"ItemOpTime"`
}

type NotifyNodes struct {
	NotifyNode []NotifyNode `xml:"NotifyNode" json:"NotifyNode"`
}

type NotifyNode struct {
	ItemName   string `xml:"ItemName" json:"ItemName"`
	ItemUserId string `xml:"ItemUserId" json:"ItemUserId"`
	ItemImage  string `xml:"ItemImage" json:"ItemImage"`
}

type ApprovalNodes struct {
	ApprovalNode []ApprovalNode `xml:"ApprovalNode" json:"ApprovalNode"`
}

type ApprovalInfo struct {
	Text              string      `xml:",chardata" json:",chardata"`
	SpNO              string      `xml:"SpNo" json:"SpNo"`
	SpName            string      `xml:"SpName" json:"SpName"`
	SpStatus          string      `xml:"SpStatus" json:"SpStatus"`
	TemplateID        string      `xml:"TemplateId" json:"TemplateId"`
	ApplyTime         string      `xml:"ApplyTime" json:"ApplyTime"`
	Applier           *Applier    `xml:"Applyer" json:"Applyer"`
	SpRecord          []*SPRecord `xml:"SpRecord" json:"SpRecord"`
	Notifier          *Notifier   `xml:"Notifyer" json:"Notifyer"`
	Comments          *Comments   `xml:"Comments" json:"Comments"`
	StatusChangeEvent string      `xml:"StatuChangeEvent" json:"StatuChangeEvent"`

	// new fields
	// https://developer.work.weixin.qq.com/document/path/90240#%E5%AE%A1%E6%89%B9%E7%8A%B6%E6%80%81%E9%80%9A%E7%9F%A5%E4%BA%8B%E4%BB%B6
	ThirdNo        string        `xml:"ThirdNo" json:"ThirdNo"`
	OpenSpName     string        `xml:"OpenSpName" json:"OpenSpName"`
	OpenTemplateId string        `xml:"OpenTemplateId" json:"OpenTemplateId"`
	OpenSpStatus   int           `xml:"OpenSpStatus" json:"OpenSpStatus"`
	ApplyUserName  string        `xml:"ApplyUserName" json:"ApplyUserName"`
	ApplyUserId    string        `xml:"ApplyUserId" json:"ApplyUserId"`
	ApplyUserParty string        `xml:"ApplyUserParty" json:"ApplyUserParty"`
	ApplyUserImage string        `xml:"ApplyUserImage" json:"ApplyUserImage"`
	ApprovalNodes  ApprovalNodes `xml:"ApprovalNodes" json:"ApprovalNodes"`
	NotifyNodes    NotifyNodes   `xml:"NotifyNodes" json:"NotifyNodes"`
	ApproverStep   int           `xml:"approverstep" json:"approverstep"`
}

type EventOpenApprovalChange struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID      string        `xml:"AgentID" json:"AgentID"`
	ApprovalInfo *ApprovalInfo `xml:"ApprovalInfo" json:"ApprovalInfo"`
}

type EventShareAgentChange struct {
	contract.EventInterface
	models.CallbackMessageHeader
	AgentID string `xml:"AgentID" json:"AgentID"`
}

// ----------------------------------------------------------------------------

type OptionID struct {
	Text     string   `xml:",chardata" json:",chardata"`
	OptionID []string `xml:"OptionId" json:"OptionId"`
}

type SelectItem struct {
	Text        string    `xml:",chardata" json:",chardata"`
	QuestionKey string    `xml:"QuestionKey" json:"QuestionKey"`
	OptionIDs   *OptionID `xml:"OptionIds" json:"OptionIds"`
}

type SelectItems struct {
	Text         string        `xml:",chardata" json:",chardata"`
	SelectedItem []*SelectItem `xml:"SelectedItem" json:"SelectedItem"`
}

type EventTemplateCardEvent struct {
	contract.EventInterface
	models.CallbackMessageHeader
	EventKey      string       `xml:"EventKey" json:"EventKey"`
	TaskID        string       `xml:"TaskId" json:"TaskId"`
	CardType      string       `xml:"CardType" json:"CardType"`
	ResponseCode  string       `xml:"ResponseCode" json:"ResponseCode"`
	AgentID       string       `xml:"AgentID" json:"AgentID"`
	SelectedItems *SelectItems `xml:"SelectedItems" json:"SelectedItems"`
}

type EventTemplateCardMenuEvent struct {
	contract.EventInterface
	models.CallbackMessageHeader
	// EventKey     string `xml:"EventKey" json:"EventKey"`
	TaskID       string `xml:"TaskId" json:"TaskId"`
	CardType     string `xml:"CardType" json:"CardType"`
	ResponseCode string `xml:"ResponseCode" json:"ResponseCode"`
	AgentID      string `xml:"AgentID" json:"AgentID"`
}

type EventKFMsgOrEvent struct {
	contract.EventInterface
	models.CallbackMessageHeader
	Token    string `xml:"Token" json:"Token"`
	OpenKfID string `xml:"OpenKfId" json:"OpenKfId"`
}
