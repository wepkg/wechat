package message

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
