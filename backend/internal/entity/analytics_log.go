package entity

type AnalyticsLog struct {
	BaseEntity

	LogType      string  `json:"log_type"`
	UserID       *string `json:"user_id,omitempty"`
	RoomID       *string `json:"room_id,omitempty"`
	Content      string  `json:"content"`
	IsToxic      bool    `json:"is_toxic"`
	AIDiagnostic string  `json:"ai_diagnostic,omitempty"`
}
