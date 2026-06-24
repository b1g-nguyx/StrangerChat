package entity

type Report struct {
	BaseEntity
	ReporterID string `json:"reporter_id" db:"reporter_id"`
	ReportedID string `json:"reported_id" db:"reported_id"`
	Reason     string `json:"reason" db:"reason"`
	Status     string `json:"status" db:"status"` // e.g., "PENDING", "RESOLVED", "DISMISSED"
}
