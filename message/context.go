package message

import "encoding/xml"

// Context ..
type Context struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
}

// Message ..
// type Message interface{}

// TextMessage ..
type TextMessage struct {
	Content string
	MsgId   int64
}

// TextMessage ..
type TextMessage struct {
	Content string
	MsgId   int64
}

// EventMessage ..
type EventMessage struct {
	Event    string
	EventKey string
	Ticket   string
}

// LocationEventMessage ..
type LocationEventMessage struct {
	Latitude  int64
	Longitude int64
	Precision int64
}

// Msg ..
type Msg struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	MsgEncrypt
	MsgPlain
}

// MsgEncrypt ..
type MsgEncrypt struct {
	Encrypt string
}

// MsgPlain ..
type MsgPlain struct {
	URL          string `xml:"URL"`
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

// <xml>
// <ToUserName><![CDATA[gh_4586052027b9]]></ToUserName>
// <FromUserName><![CDATA[o8UFh1io3RWBWnW3jT4CXmHB8X2k]]></FromUserName>
// <CreateTime>1608701921</CreateTime>
// <MsgType><![CDATA[text]]></MsgType>
// <Content><![CDATA[外卖返利]]></Content>
// <MsgId>23031074913322685</MsgId>
// <Encrypt><![CDATA[TJdJn4637WrAq7laOtbVO1ocIhTuT01D0llNzMv40Ge1NBh1OZaf905MP6eQe+0bI6ZAaZTaLloxi2eLy6C1Zn0BhiqJXRpp91jLxj516D/EQRceHQfOWesgLJIlKTCLJPEuszDT55CpfYfCAgogvIFXpBZDr0TKHbewa3WoetWP2Wtc4jXzbfyR0um9COtWnEZ6HuBOKM/AKPu1QlMJl9zzq5LETs3POhFuJgOoKe3AjF8pqSSI4Ty6s8vWbrcKKJ8FPnFnSwhraDWI9kCe7slxQrddRTl4aMTZO+aLkAJp2M9M0BmibdW373hSx0Rb98GGHMDRk58sNkJeSZueNIiRDUJUJN/qpueEVyVZJyZnBQSLzIyL06bBZJtqdKQY7QBQxEPl2HuarJiHH6BjXvkVvLFtzoCaAr8Q9SjLcQnwSW/hWa1M6QBCLI/4t6KLTKNdl1JDyGq0ezedv7r46Q==]]></Encrypt>
// </xml>
