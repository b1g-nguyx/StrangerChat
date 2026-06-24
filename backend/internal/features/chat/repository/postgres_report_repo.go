package repository

import (
	"context"

	"github.com/b1g-nguyx/strangerchat-backend/internal/core/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/repo/persistent"
)

type PostgresReportRepo interface {
	CreateReport(ctx context.Context, report entity.Report) error
	SaveEvidence(ctx context.Context, evidence entity.ReportEvidence) error
}

type postgresReportRepoImpl struct {
	persistent.BaseRepo
}

func NewPostgresReportRepo(base persistent.BaseRepo) PostgresReportRepo {
	return &postgresReportRepoImpl{BaseRepo: base}
}

func (r *postgresReportRepoImpl) CreateReport(ctx context.Context, report entity.Report) error {
	query, args, err := r.Builder.Insert("reports").
		Columns("id", "reporter_id", "reported_id", "reason", "status", "created_at", "updated_at").
		Values(report.ID, report.ReporterID, report.ReportedID, report.Reason, report.Status, report.CreatedAt, report.UpdatedAt).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, query, args...)
	return err
}

func (r *postgresReportRepoImpl) SaveEvidence(ctx context.Context, evidence entity.ReportEvidence) error {
	query, args, err := r.Builder.Insert("report_evidences").
		Columns("id", "report_id", "room_id", "chat_logs", "created_at", "updated_at").
		Values(evidence.ID, evidence.ReportID, evidence.RoomID, evidence.ChatLogs, evidence.CreatedAt, evidence.UpdatedAt).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, query, args...)
	return err
}
