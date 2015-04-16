package event

import (
	// "github.com/astaxie/beego"
	"encoding/xml"
)

type ViewEvent struct {
	Id           int64
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	EventKey     string `xml:"EventKey"`
}

func ReceiveViewEvent(content string) string {
	var msg ViewEvent
	err := xml.Unmarshal([]byte(content), &msg)
	if err != nil {
		return ""
	}
	return ""
}
