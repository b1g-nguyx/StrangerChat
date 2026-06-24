package entity

type ReportEvidence struct {
	BaseEntity
	ReportID string `json:"report_id" db:"report_id"`
	RoomID   string `json:"room_id" db:"room_id"`
	ChatLogs []byte `json:"chat_logs" db:"chat_logs"` // JSONB in Postgres
}
