package messages

import (
	"github.com/uvlydxba56648/PowerWeChat/v3/src/kernel/power"
)

type Image struct {
	*Media
}

func NewImage(mediaID string, items *power.HashMap) *Image {
	m := &Image{
		NewMedia(mediaID, "image", items),
	}

	return m
}
