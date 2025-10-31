package oa

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/approval"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/calendar"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/dial"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/journal"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/living"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/meeting"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/meetingroom"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/pstncc"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/schedule"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/wedoc"
	"github.com/uvlydxba56648/PowerWeChat/v3/src/work/oa/wedrive"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client,
	*calendar.Client,
	*dial.Client,
	*journal.Client,
	*living.Client,
	*meeting.Client,
	*meetingroom.Client,
	*pstncc.Client,
	*schedule.Client,
	*wedoc.Client,
	*wedrive.Client,
	*approval.Client,
	error,

) {
	//config := app.GetConfig()

	Client, err := NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Calendar, err := calendar.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Dial, err := dial.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Journal, err := journal.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Living, err := living.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Meeting, err := meeting.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	MeetingRoom, err := meetingroom.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	PSTNCC, err := pstncc.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Schedule, err := schedule.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	WeDoc, err := wedoc.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	WeDrive, err := wedrive.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	Approval, err := approval.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	return Client,
		Calendar,
		Dial,
		Journal,
		Living,
		Meeting,
		MeetingRoom,
		PSTNCC,
		Schedule,
		WeDoc,
		WeDrive,
		Approval,
		nil

}
