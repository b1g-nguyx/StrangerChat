ALTER TABLE "report_evidences" DROP CONSTRAINT IF EXISTS "report_evidences_report_id_fkey";
ALTER TABLE "report_evidences" DROP CONSTRAINT IF EXISTS "report_evidences_room_id_fkey";
ALTER TABLE "reports" DROP CONSTRAINT IF EXISTS "reports_reporter_id_fkey";
ALTER TABLE "reports" DROP CONSTRAINT IF EXISTS "reports_reported_id_fkey";

DROP TABLE IF EXISTS "report_evidences";
DROP TABLE IF EXISTS "reports";
