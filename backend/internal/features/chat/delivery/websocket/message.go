package websocket

type WSMessageType string

const (
	MsgTypeFindMatch   WSMessageType = "FIND_MATCH"
	MsgTypeMatched     WSMessageType = "MATCHED"
	MsgTypeChat        WSMessageType = "CHAT"
	MsgTypeLeave       WSMessageType = "LEAVE_ROOM"
	MsgTypePartnerLeft WSMessageType = "PARTNER_LEFT"
	MsgTypeReport      WSMessageType = "REPORT"
)

type WSMessage struct {
	Type       WSMessageType `json:"type"`
	RoomID     string        `json:"room_id,omitempty"`
	ReportedID string        `json:"reported_id,omitempty"`
	Content    string        `json:"content,omitempty"`
}

type BroadcastMessage struct {
	RoomID  string
	Message []byte
}
